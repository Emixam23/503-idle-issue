package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	s := &http.Server{
		Addr:              ":3000",
		Handler:           mux,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	log.Println("Starting server on :3000")
	log.Fatal(s.ListenAndServe())
}
