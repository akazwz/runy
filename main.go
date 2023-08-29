package main

import (
	"encoding/json"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"message": "Hello World",
		}
		JSON(w, data, http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)
}

func JSON(w http.ResponseWriter, data any, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}
