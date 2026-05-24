package store

import (
	"encoding/json"
	"os"

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
