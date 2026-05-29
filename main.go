package main

import (
	"log"
	"net/http"

	"github.com/recursion-gowebapi/go-web-api/handlers"
)

func withCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("web")))

	http.HandleFunc("/api/slangs", handlers.SlangsHandler)
	http.HandleFunc("/api/slangs/search", handlers.SearchSlangsHandler)
	http.HandleFunc("/api/slangs/{id}", handlers.DetailSlangHandler)
	http.HandleFunc("/api/slangs/filter", handlers.Filter)
	http.HandleFunc("/api/slangs/random", handlers.RandomSlangsHandler)
	http.HandleFunc("/api/slangs/add", handlers.CreateSlang)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", withCORS(http.DefaultServeMux)))
}
