package handlers

import (
	"encoding/json"
	"net/http"
	"fmt"

	//"github.com/tutorials/go/crud/pkg/mocks"
	"github.com/ThotPrime/Project/tree/master/Project/pkg/models"
	//"github.com/ThotPrime/Project/tree/main/Project/pkg/mocks"
)

func (h handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book

	if result := h.DB.Find(&books); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}