package store

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/recursion-gowebapi/go-web-api/models"
)

func GetSlangs() ([]models.Slang, error) {
	// データ取得
	data, err := os.ReadFile("data/slangs.json")
	if err != nil {
		return nil, err
	}
	// データを構造体に変換
	var slangs []models.Slang
	if err := json.Unmarshal(data, &slangs); err != nil {
		return nil, err
	}
	return slangs, nil
}

// stringの小文字化, 空白削除用関数
func normalizeText(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

func SearchSlangs(keyword string) ([]models.Slang, error) {
	// データ取得
	slangs, err := GetSlangs()
	if err != nil {
		return nil, err
	}

	// keywordを小文字化, 空白削除
	normalizedKeyword := normalizeText(keyword)

	results := []models.Slang{}

	// slangを小文字化, 空白削除し検索
	for _, slang := range slangs {
		normalizedSlang := normalizeText(slang.Slang)

		if normalizedSlang == normalizedKeyword {
			results = append(results, slang)
		}
	}

	return results, nil
}
