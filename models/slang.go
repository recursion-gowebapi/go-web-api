package models

type Slang struct {
	ID       string    `json:"id"`
	Slang    string    `json:"slang"`
	Meanings []Meaning `json:"meanings"`
}

type Meaning struct {
	MeaningJa         string    `json:"meaning_ja"`
	NuanceJa          string    `json:"nuance_ja"`
	Examples          []Example `json:"examples"`
	Scene             []string  `json:"scene"`
	EmotionCategories []string  `json:"emotion_categories"`
	WarningJa         string    `json:"warning_ja"`
}

type Example struct {
	SentenceEn string `json:"sentence_en"`
	SentenceJa string `json:"sentence_ja"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
