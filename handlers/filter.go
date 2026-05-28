package handlers

import (
	"encoding/json"
	"net/http"
	"slices"

	"github.com/recursion-gowebapi/go-web-api/models"
	"github.com/recursion-gowebapi/go-web-api/store"
)

func Filter(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	var result []models.Slang
	//JSONファイルの読み込み

	slangs, err := store.GetSlangs()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	//クエリの読み込み
	scene := r.URL.Query().Get("scene")
	emotion := r.URL.Query().Get("emotion")

	if scene == "" && emotion == "" {
		respondWithError(w, http.StatusBadRequest, "Invalid parameter")
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
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resultJSON)
}

//そのスラングにワードが含まれているか確認し、bool値を返す
func isContain(slang models.Slang, targetWord string, category string) bool {
	for _, meaning := range slang.Meanings {
		var target []string

		switch category {
		case "emotion":
			target = meaning.EmotionCategories
		case "scene":
			target = meaning.Scene
		}

		if slices.Contains(target, targetWord) {
			return true
		}
	}
	return false
}
