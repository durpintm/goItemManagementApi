package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Item struct represents an item with ID and Name
type ItemTest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var testItems = []ItemTest{} // In-memory storage for items

func TestGreetHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/items", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create a new router and handle the request
	router := mux.NewRouter()
	router.HandleFunc("/items", getItemsHandler).Methods("GET")
	router.ServeHTTP(rr, req)
	// Check the status code of the response
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("The GET [getItemsHandler] returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the content type of the response
	if contentType := rr.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("The GET [getItemsHandler] returned unexpected content type: got %v want %v", contentType, "application/json")
	}

}

func TestAddItemHandler(t *testing.T) {
	// Create a new item to add
	newItem := Item{
		Name: "Laptop",
		ID:   uuid.New().String(),
	}
	jsonData, err := json.Marshal(newItem)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new request to add the item
	req, err := http.NewRequest("POST", "/items/create", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

}
