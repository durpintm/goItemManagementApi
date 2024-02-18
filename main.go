package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Item struct represents an item with ID and Name
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{} // In-memory storage for items
const Dport = ":8000"

func main() {
	fmt.Println(items)
	http.HandleFunc("/items", getItemsHandler)

	fmt.Printf("Server is starting on port: %v\n", Dport) // Added newline for better terminal output
	http.ListenAndServe(Dport, nil)
}

// this handler retrieves a list of all items
func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
