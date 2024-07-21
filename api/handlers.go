package api

import (
	"encoding/json"
	"net/http"
	"stockx-monitor/internal/monitor"
)

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
    products, err := monitor.GetProductDetails()
    if err != nil {
        http.Error(w, "Failed to get product details: "+err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(products)
}

