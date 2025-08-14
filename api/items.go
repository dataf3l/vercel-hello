package handler

import (
	"encoding/json"
	"net/http"
)

// Item represents a simple data structure
type Item struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
}

// Handler returns a JSON list of 5 hardcoded items
func Handler(w http.ResponseWriter, r *http.Request) {
	// Set content type to JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// Handle preflight requests
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Create 5 hardcoded items
	items := []Item{
		{
			ID:          1,
			Name:        "Wireless Headphones",
			Description: "High-quality wireless headphones with noise cancellation",
			Category:    "Electronics",
			Price:       199.99,
		},
		{
			ID:          2,
			Name:        "Coffee Mug",
			Description: "Ceramic coffee mug with heat retention technology",
			Category:    "Kitchen",
			Price:       24.99,
		},
		{
			ID:          3,
			Name:        "Notebook",
			Description: "Premium leather-bound notebook with lined pages",
			Category:    "Stationery",
			Price:       39.99,
		},
		{
			ID:          4,
			Name:        "Smartphone Case",
			Description: "Protective case with wireless charging compatibility",
			Category:    "Accessories",
			Price:       29.99,
		},
		{
			ID:          5,
			Name:        "Desk Lamp",
			Description: "LED desk lamp with adjustable brightness and color temperature",
			Category:    "Furniture",
			Price:       79.99,
		},
	}

	// Encode and send the JSON response
	if err := json.NewEncoder(w).Encode(items); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}
}
