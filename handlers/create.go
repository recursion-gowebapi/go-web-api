package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/recursion-gowebapi/go-web-api/models"
	"github.com/recursion-gowebapi/go-web-api/store"
)

func CreateSlang(w http.ResponseWriter, r *http.Request) {

	//メソッドチェック
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	//スラング取得
	slangs, err := store.GetSlangs()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	var newSlang models.Slang
	//jsonをデコード
	err = json.NewDecoder(r.Body).Decode(&newSlang)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	//slangのチェック
	if len(newSlang.Slang) == 0 {
		respondWithError(w, http.StatusBadRequest, "slang is required")
		return
	}
	//idのチェック
	if len(newSlang.ID) == 0 {
		respondWithError(w, http.StatusBadRequest, "id is required")
		return
	}
	//meaningのチェック
	if len(newSlang.Meanings) == 0 {
		respondWithError(w, http.StatusBadRequest, "meanings is required")
		return
	}

	isContainSlang, err := store.SearchSlangs(newSlang.Slang)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	isContainID, err := store.GetSlangByID(newSlang.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	//slangがすでにあるか
	if len(isContainSlang) != 0 {
		respondWithError(w, http.StatusBadRequest, "The slang already exist")
		return
	}

	//IDが重複していないか
	if isContainID != nil {
		respondWithError(w, http.StatusBadRequest, "The ID already exist")
		return
	}

	//slangsに追加
	slangs = append(slangs, newSlang)

	//ファイルの書き換え
	err = store.Save(slangs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to save data")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSlang)
}
