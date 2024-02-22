package main

import (
	"encoding/json"
	"net/http"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		items := []Item{
			{ID: 1, Name: "Item 1"},
			{ID: 2, Name: "Item 2"},
			{ID: 3, Name: "Item 3"},
		}
		json.NewEncoder(w).Encode(items)
	})

	http.ListenAndServe(":8080", nil)
}
