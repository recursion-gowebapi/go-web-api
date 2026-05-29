package store

import (
	"encoding/json"
	"math/rand"
	"os"
	"fmt"
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

func GetSlangByID(id string) (*models.Slang, error) {
	slangs, err := GetSlangs()
	if err != nil {
		return nil, err
	}
	for _, slang := range slangs {
		if slang.ID == id {
			return &slang, nil
		}
	}
	return nil, nil
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

	for _, slang := range slangs {
		if slang.ID == id {
			return &slang, nil
		}
	}

	return nil, fmt.Errorf("slang not found")
}

func SearchSlangs(keyword string) ([]models.Slang, error) {
  slangs, err := GetSlangs()
  
  if err != nil {
    return nil, err
  }
  
  normalizedKeyword := normalizeText(keyword)

	results := []models.Slang{}

	for _, slang := range slangs {
		normalizedSlang := normalizeText(slang.Slang)

		if normalizedSlang == normalizedKeyword {
			results = append(results, slang)
		}
	}

	return results, nil
}

func RandomSlangs(count int) ([]models.Slang, error) {
	// データ取得
	slangs, err := GetSlangs()
	if err != nil {
		return nil, err
	}

	// countの上限をスラング配列の長さに設定
	if count > len(slangs) {
		count = len(slangs)
	}

	// スラング配列をシャッフル
	rand.Shuffle(len(slangs), func(i, j int) {
		slangs[i], slangs[j] = slangs[j], slangs[i]
	})

	// count件取得
	results := slangs[:count]

	return results, nil
}

//jsonファイルの書き換え
func Save(slangs []models.Slang) error {
	data, err := json.MarshalIndent(slangs, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile("data/slangs.json", data, 0644)
}
