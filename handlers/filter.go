package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/recursion-gowebapi/go-web-api/models"
)

func filter(w http.ResponseWriter, r *http.Request) {
	var result []models.Slang
	//JSONファイルの読み込み
	jsonFile, err := os.Open("slang.json")
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	defer jsonFile.Close()
	jsonData, err := io.ReadAll(jsonFile)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	var slangs []models.Slang
	err = json.Unmarshal(jsonData, &slangs)
	if err != nil {
		writeError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	//クエリの読み込み
	scene := r.URL.Query().Get("scene")
	emotion := r.URL.Query().Get("emotion")

	if scene == "" && emotion == "" {
		writeError(w, http.StatusBadRequest, "Invalid parameter")
		return
	}

	//当てはまるslangをappend
	for _, slang := range slangs {
		if scene == "" {
			if isContain(slang, emotion, "emotion") {
				result = append(result, slang)
			}
		} else if emotion == "" {
			if isContain(slang, scene, "scene") {
				result = append(result, slang)
			}
		} else {
			if isContain(slang, scene, "scene") && isContain(slang, emotion, "emotion") {
				result = append(result, slang)
			}
		}
	}

	//json形式にエンコードし、結果を返す
	resultJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

//そのスラングにワードが含まれているか確認し、bool値を返す
func isContain(slang models.Slang, word string, category string) bool {
	for _, meaning := range slang.Meanings {

		if category == "emotion" {
			for _, words := range meaning.EmotionCategories {
				if words == word {
					return true
				}
			}
		} else if category == "scene" {
			for _, words := range meaning.Scene {
				if words == word {
					return true
				}
			}
		}
	}
	return false
}

//エラーが生じた場合、エラーを返す
func writeError(w http.ResponseWriter, status int, msg string) {
	res := models.ErrorResponse{
		Error: msg,
	}
	resJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, "encode error failed", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(resJSON)
}
