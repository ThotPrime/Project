package handlers

import (
	"encoding/json"
	"net/http"

	//"github.com/tutorials/go/crud/pkg/mocks"
	"github.com/ThotPrime/Project/tree/main/Project/pkg/mocks"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mocks.Books)
}