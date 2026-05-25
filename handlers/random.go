package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/recursion-gowebapi/go-web-api/models"
	"github.com/recursion-gowebapi/go-web-api/store"
)

type RandomResponse struct {
	Status string         `json:"status"`
	Count  int            `json:"count"`
	Items  []models.Slang `json:"items"`
}

// GET /api/slangs/random?count={number}
func RandomSlangsHandler(w http.ResponseWriter, r *http.Request) {

	// メソッドチェック後ほど変更予定
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// レスポンス形式をJSONに指定
	w.Header().Set("Content-Type", "application/json")

	// countのデフォルト値を設定
	count := 1

	// countを取得
	countParam := r.URL.Query().Get("count")

	// countが指定されている場合、文字列から整数に変換
	if countParam != "" {
		parsedCount, err := strconv.Atoi(countParam)
		if err != nil || parsedCount <= 0 {
			respondWithError(w, http.StatusBadRequest, "count must be a positive integer")
			return
		}

		count = parsedCount
	}

	// ランダムスラングを取得
	results, err := store.RandomSlangs(count)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	// ランダム取得結果をJSONで返す
	response := RandomResponse{
		Status: "success",
		Count:  len(results),
		Items:  results,
	}

	json.NewEncoder(w).Encode(response)
}
