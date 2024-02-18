package main

import "fmt"

// Item struct represents an item with ID and Name
type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items = []Item{} // In-memory storage for items
const Dport = ":8000"

func main() {
	fmt.Println(items)
}