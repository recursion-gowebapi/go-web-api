package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/recursion-gowebapi/go-web-api/models"
	"github.com/recursion-gowebapi/go-web-api/store"
)

type SearchResponse struct {
	Status string         `json:"status"`
	Count  int            `json:"count"`
	Items  []models.Slang `json:"items"`
}

// PR#9 marge後に削除 19-21行
type ErrorResponse struct {
	Error string `json:"error"`
}

// GET /api/slangs/search?keyword={word}
func SearchSlangsHandler(w http.ResponseWriter, r *http.Request) {

	//メソッドチェック
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	// レスポンス形式をJSONに指定
	w.Header().Set("Content-Type", "application/json")

	// keywordを取得
	keyword := r.URL.Query().Get("keyword")

	// keyword未指定チェック
	if strings.TrimSpace(keyword) == "" {
		w.WriteHeader(http.StatusBadRequest)
		//PR#9 marge後にmodels.ErrorResponseに置き換え
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "keyword parameter is required",
		})
		return
	}

	// keywordでスラング検索
	results, err := store.SearchSlangs(keyword)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//PR#9 marge後にmodels.ErrorResponseに置き換え
		json.NewEncoder(w).Encode(ErrorResponse{
			Error: "internal server error",
		})
		return
	}

	// 検索結果をJSONで返す
	response := SearchResponse{
		Status: "success",
		Count:  len(results),
		Items:  results,
	}

	json.NewEncoder(w).Encode(response)
}

//メモ：３箇所PR#9 marge後に変更箇所あり
