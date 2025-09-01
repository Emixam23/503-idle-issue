package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const (
	timeout = 10 * time.Second
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err := fmt.Fprintln(w, "Hello, world!")
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	r := mux.NewRouter()

	// GET /
	r.HandleFunc("/", helloWorldHandler).Methods(http.MethodGet)

	muxWithMiddlewares := http.TimeoutHandler(r, timeout, "Timeout!")

	addr := ":8080"
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, muxWithMiddlewares))
}
