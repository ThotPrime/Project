package main

import (
	"log"
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
	"github.com/ThotPrime/Project/tree/master/Project/pkg/db"
	//"github.com/Project/go/crud/pkg/handlers"
	//"github.com/karanpratapsingh/tutorials/tree/master/go/gorm/pkg/db"
  	"github.com/ThotPrime/Project/tree/master/Project/pkg/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)


var DBM *mongo.Collection

func main() {
	DB := db.InitP()
	h := handlers.New(DB)
	router := mux.NewRouter()

	DBM := db.initMongo()
	hM := handlers.New(DBM)
	r := mux.NewRouter()
	

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet)
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet)
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost)
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut)
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)
	router.HandleFunc("/api/movies", hM.GetMyAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", hM.CreateMovie).Methods("POST")

	
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4001", r))
	fmt.Println("Listening at port 4001 ...")

	log.Println("API is running!")
	http.ListenAndServe(":4000", router)
}