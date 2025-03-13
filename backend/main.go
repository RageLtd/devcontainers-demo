package main

import (
	"fmt"
	"net/http"
)

func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
        // Set CORS headers
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

        if r.Method == "OPTIONS" {
            return
        }

        fmt.Fprint(w, "{ \"message\": \"Hello!\" }")
    })

    fmt.Println("Server starting on :8080...")
    http.ListenAndServe("0.0.0.0:8080", nil)
}
