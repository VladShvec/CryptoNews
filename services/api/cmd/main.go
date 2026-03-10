package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		_ = json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"service": "api",
		})
	})

	//mux.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(http.StatusOK)
	//
	//	_ = json.NewEncoder(w).Encode(map[string]any{
	//		"data": []map[string]any{
	//			{
	//				"id":    1,
	//				"title": "Bitcoin jumps after ETF inflows",
	//			},
	//			{
	//				"id":    2,
	//				"title": "Ethereum ecosystem sees renewed activity",
	//			},
	//		},
	//	})
	//})

	log.Println("api service started on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
