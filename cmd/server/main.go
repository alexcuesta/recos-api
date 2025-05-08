package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/recommendations", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Recommendations"))
	})

	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
