package handler

import (
	"encoding/json"
	"net/http"
)

// Person structure
type Person struct {
	Name string
	Age  int
}

// Get handler returns a Person
func Get(w http.ResponseWriter, r *http.Request) {
	p := Person{
		"angel",
		23,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
