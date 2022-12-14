package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"net/http"
	"fmt"

	//"github.com/tutorials/go/crud/pkg/mocks"
	//"github.com/ThotPrime/Project/tree/main/Project/pkg/mocks"
	//"github.com/tutorials/go/crud/pkg/models"
	"github.com/ThotPrime/Project/tree/master/Project/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}
	
	var book models.Book
	json.Unmarshal(body, &book)

	// Append to the Books table
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}