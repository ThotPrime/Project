package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ThotPrime/Project/tree/master/pkg/db"
	//"github.com/Project/go/crud/pkg/handlers"
	//"github.com/karanpratapsingh/tutorials/tree/master/go/gorm/pkg/db"
  	"github.com/ThotPrime/Project/tree/main/Project/pkg/handlers"
)

func main() {
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}