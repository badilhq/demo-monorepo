package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("encode response: %v", err)
	}
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]string{"status": "ok", "service": "api"})
	})
	mux.HandleFunc("GET /hello", func(w http.ResponseWriter, _ *http.Request) {
		writeJSON(w, map[string]string{"service": "api", "message": "hello from api"})
	})
	return mux
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("api listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, newMux()))
}
// e2e 20261128101920
