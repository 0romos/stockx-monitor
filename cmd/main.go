package main

import (
	"log"
	"net/http"
	"stockx-monitor/api"
)

func main() {
	http.HandleFunc("/products", api.ProductsHandler)

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
