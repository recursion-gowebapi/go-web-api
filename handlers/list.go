package handlers

import (
	"encoding/json"
	"github.com/recursion-gowebapi/go-web-api/models"
	"github.com/recursion-gowebapi/go-web-api/store"
	"net/http"
	"strconv"
)

// GET /api/slangs
func SlangsHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	queryMap := r.URL.Query()

	limit := 20
	if l, ok := queryMap["limit"]; ok && len(l) > 0 {

		var err error
		limit, err = strconv.Atoi(l[0])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "invalid limit parameter")
			return
		}
		if limit < 1 || limit > 100 {
			respondWithError(w, http.StatusBadRequest, "invalid limit parameter")
			return
		}
	}

	offset := 0
	if o, ok := queryMap["offset"]; ok && len(o) > 0 {
		var err error
		offset, err = strconv.Atoi(o[0])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "invalid offset parameter")
			return
		}
		if offset < 0 {
			respondWithError(w, http.StatusBadRequest, "invalid offset parameter")
			return
		}
	}

	slangs, err := store.GetSlangs()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	// offset がデータの総件数を超えた場合、空配列を返す
	if offset > len(slangs) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("[]"))
		return
	}

	var resSlangs []models.Slang
	if offset+limit > len(slangs) {
		resSlangs = slangs[offset:]
	} else {
		resSlangs = slangs[offset : offset+limit]
	}

	jsonData, err := json.Marshal(resSlangs)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func respondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
