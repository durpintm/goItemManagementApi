package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Item struct represents an item with ID and Name
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{} // In-memory storage for items
const Dport = ":8000"

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/items/create", addItemHandler).Methods("POST")
	router.HandleFunc("/items", getItemsHandler).Methods("GET")

	// http.HandleFunc("/items", getItemsHandler)
	// http.HandleFunc("/items/create", addItemHandler)

	fmt.Printf("Server is starting on port: %v\n", Dport) // Added newline for better terminal output
	http.ListenAndServe(Dport, router)
		// http.ListenAndServe(Dport, nil)
}

// this handler retrieves a list of all items
func getItemsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//For testing purpose only -- uncomment to test
	// items = append(items, Item{ID: uuid.New().String(), Name: "Item 1"})
  // items = append(items, Item{ID: uuid.New().String(), Name: "Item 2"})	
	json.NewEncoder(w).Encode(items)
}

// this handler adds a new item to the collection
func addItemHandler(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generate an unique ID for the new item 
	newItem.ID = uuid.New().String()
	// newItem.Name = "Laptop" //For testing
	items = append(items, newItem)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newItem)
}