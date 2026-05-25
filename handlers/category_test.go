package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"reflect"
	"github.com/recursion-gowebapi/go-web-api/models"
)

func TestMethodCheck_OK(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/categories/scenes",
		nil,
	)

	rr := httptest.NewRecorder()

	result := methodCheck(rr, req, http.MethodGet)

	if !result {
		t.Error("expected true, got false")
	}
}

func TestMethodCheck_NG(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/categories/scenes",
		nil,
	)

	rr := httptest.NewRecorder()

	result := methodCheck(rr, req, http.MethodGet)

	if result {
		t.Error("expected false, got true")
	}

	if rr.Code != http.StatusMethodNotAllowed {
		t.Errorf("expected 405, got %d", rr.Code)
	}
}

func TestValidCategoryType_OK(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/categories/scenes",
		nil,
	)

	rr := httptest.NewRecorder()

	validParams := []string{"scenes", "emotion_categories"}

	param, err := validateCategoryType(rr, req, validParams)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if param != "scenes" {
		t.Errorf("expected 'scenes', got '%s'", param)
	}
}

func TestValidCategoryType_NG(t *testing.T) {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/categories/invalid",
		nil,
	)

	rr := httptest.NewRecorder()

	validParams := []string{"scenes", "emotion_categories"}

	param, err := validateCategoryType(rr, req, validParams)

	if err == nil {
		t.Error("expected error, got nil")
	}

	if param != "" {
		t.Errorf("expected empty string, got '%s'", param)
	}
}

func TestUniqueCategories_Empty(t *testing.T) {
	slangs := []models.Slang{}

	var expected = []string{}
	result := getUniqueCategories(slangs, "scenes")
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGetUniqueCategories_OK(t *testing.T) {
	slangs := []models.Slang {
		{
			ID: "slang_001", 
			Slang: "lit", 
			Meanings: []models.Meaning {
				{
					MeaningJa: "...", 
					Scene: []string{"friends", "family", "party", "work", "school"}, 
					EmotionCategories: []string{"joy", "surprise", "neutral", "other", "unknown"},
				},
			},
		}, 
		{
			ID: "slang_002", 
			Slang: "salty", 
			Meanings: []models.Meaning {
				{
					MeaningJa: "...",
					Scene: []string{"family", "work", "school", "other", "unknown"},
					EmotionCategories: []string{"anger", "disgust", "neutral", "other", "unknown"},
				},
			},
		 },
	}

	expectedScenes := []string{"family", "friends", "other", "party", "school", "unknown", "work"}
	expectedEmotions := []string{"anger", "disgust", "joy", "neutral", "other", "surprise", "unknown"}

	scenes := getUniqueCategories(slangs, "scenes")
	emotions := getUniqueCategories(slangs, "emotion_categories")

	if !reflect.DeepEqual(scenes, expectedScenes) {
		t.Errorf("expected %v, got %v", expectedScenes, scenes)
	}

	if !reflect.DeepEqual(emotions, expectedEmotions) {
		t.Errorf("expected %v, got %v", expectedEmotions, emotions)
	}
}
