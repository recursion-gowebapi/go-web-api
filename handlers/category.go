package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/recursion-gowebapi/go-web-api/models"
	"github.com/recursion-gowebapi/go-web-api/store"
	"net/http"
	"sort"
	"strings"
)

// 共通のメソッドチェック関数
func methodCheck(w http.ResponseWriter, r *http.Request, allowedMethod string) bool {
	if r.Method != allowedMethod {
		w.Header().Set("Allow", allowedMethod)
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return false
	}
	return true
}

// パラメータの確認
func validateCategoryType(w http.ResponseWriter, r *http.Request, validParams []string) (string, error) {
	path := r.URL.Path
	params := strings.Split(path, "/")
	param := params[len(params)-1]

	for _, validParam := range validParams {
		if param == validParam {
			return param, nil
		}
	}

	respondWithError(w, http.StatusBadRequest, "invalid category type")
	return "", fmt.Errorf("invalid category type")
}

func getUniqueCategories(slangs []models.Slang, categoryType string) []string {
	uniqueCategories := make(map[string]struct{})

	for _, slang := range slangs {
		for _, meaning := range slang.Meanings {
			var targets []string

			if categoryType == "scenes" {
				targets = meaning.Scene
			} else {
				targets = meaning.EmotionCategories
			}

			for _, item := range targets {
				uniqueCategories[item] = struct{}{}
			}

		}
	}

	var result []string
	for key := range uniqueCategories {
		result = append(result, key)
	}

	sort.Strings(result)
	return result
}

// GET /api/categories/{scenes or emotion_categories}
func ReturnCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	// メソッドチェック
	if !methodCheck(w, r, http.MethodGet) {
		return
	}

	// パラメータチェック
	categoryType, err := validateCategoryType(w, r, []string{"scenes", "emotion_categories"})

	if err != nil {
		return
	}

	// データ取得
	slangs, err := store.GetSlangs()

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	// カテゴリーの重複を排除して取得
	result := getUniqueCategories(slangs, categoryType)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
