package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/recursion-gowebapi/go-web-api/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Go Web API is running")
	})

	http.HandleFunc("/api/slangs", handlers.SlangsHandler)
	http.HandleFunc("/api/slangs/search", handlers.SearchSlangsHandler)
	http.HandleFunc("/api/slangs/random", handlers.RandomSlangsHandler)

	log.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
