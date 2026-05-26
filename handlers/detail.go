package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/recursion-gowebapi/go-web-api/store"
)

// GET /api/slangs/{id}
func DetailSlangHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", "GET")
		respondWithError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	id := r.PathValue("id")
	slang, err := store.GetSlangByID(id)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	if slang == nil {
		respondWithError(w, http.StatusNotFound, "slang not found")
		return
	}
	jsonData, err := json.Marshal(slang)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "internal server error")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
