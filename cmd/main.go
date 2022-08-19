package main

import (
    "log"
    "net/http"
  "encoding/json"

    "github.com/gorilla/mux"
    //"github.com/karanpratapsingh/tutorials/go/crud/pkg/handlers"
)

func main() {
    router := mux.NewRouter()

    router.HandleFunc("/Books",handlers.GetAllBooks).Methods(http.MethodGet)

    log.Println("API is running!")
    http.ListenAndServe(":4000", router)
}
