package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"fmt"

	"github.com/gorilla/mux"
	//"github.com/tutorials/go/crud/pkg/mocks"
	"github.com/ThotPrime/Project/tree/master/Project/pkg/models"
	//"github.com/ThotPrime/Project/tree/main/Project/pkg/mocks"
)

func (h handler) GetBook(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find book by Id
	var book models.Book
	
	if result := h.DB.First(&book, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}