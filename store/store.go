<<<<<<< Updated upstream
=======
package store

import (
	"encoding/json"
	"os"
	"fmt"
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

	return nil, fmt.Errorf("slang not found")
}
>>>>>>> Stashed changes
