package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	Get(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, we got: %d", http.StatusOK, w.Code)
	}

	// t.Log(w.Body.String())
	got := Person{}
	err := json.NewDecoder(w.Body).Decode(&got)
	if err != nil {
		t.Errorf("can not process json: %v", err)
	}

	want := Person{
		Name: "angel",
		Age:  23,
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("we want %v, we got %v", want, got)
	}
}
