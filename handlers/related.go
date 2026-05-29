package handlers

import (
	"encoding/json"
	"github.com/recursion-gowebapi/go-web-api/models"
	"github.com/recursion-gowebapi/go-web-api/store"
	"net/http"
	"strings"
)

// GET /api/slangs/{{id}}/related
func GetRelatedSlangs(w http.ResponseWriter, r *http.Request) {
	// メソッドチェック
	if !methodCheck(w, r, http.MethodGet) {
		return
	}

	// クエリパラメータの取得
	path := r.URL.Path
	params := strings.Split(path, "/")

	// パラメータの数が足りない場合はエラー
	if len(params) < 5 {
		respondWithError(w, http.StatusBadRequest, "invalid request path")
		return
	}

	idStr := params[len(params)-2]

	// id に対応するスラングの取得
	slang, err := store.GetSlangByID(idStr)

	if err != nil || slang == nil {
		respondWithError(w, http.StatusNotFound, "slang not found")
		return
	}

	// id が持つシーンと感情カテゴリの取得
	targetScenes := make(map[string]struct{})
	targetEmotions := make(map[string]struct{})

	for _, meaning := range slang.Meanings {
		for _, scene := range meaning.Scene {
			targetScenes[scene] = struct{}{}
		}

		for _, emotion := range meaning.EmotionCategories {
			targetEmotions[emotion] = struct{}{}
		}
	}

	// 関連スラングの取得
	// 同じシーン or 同じ感情カテゴリを持つスラングを関連スラングとする
	slangList, err := store.GetSlangs()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "failed to retrieve slangs")
		return
	}

	relatedSlangs := []models.Slang{}

	for _, s := range slangList {
		// 自分自身は関連スラングに含めない
		if s.ID == slang.ID {
			continue
		}

		matched := false

		for _, meaning := range s.Meanings {
			for _, scene := range meaning.Scene {
				// targetScenes と比較
				if _, exist := targetScenes[scene]; exist {
					matched = true
					break
				}
			}

			for _, emotion := range meaning.EmotionCategories {
				// targetEmotions と比較
				if _, exist := targetEmotions[emotion]; exist {
					matched = true
					break
				}
			}

			if matched {
				break
			}
		}

		if matched {
			relatedSlangs = append(relatedSlangs, s)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(relatedSlangs)
}
