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
		respondWithError(w, http.StatusInternalServerError, "Internal server error")
		return
	}

	//slangsに追加
	slangs = append(slangs, newSlang)

	//ファイルの書き換え
	err = store.Save(slangs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "The data saving does not work")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newSlang)
}
