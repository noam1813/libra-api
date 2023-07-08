package main

import (
	PostController "app/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", PostController.Index).Methods("GET")
	router.HandleFunc("/", PostController.Create).Methods("POST")
	router.HandleFunc("/{id}", PostController.Show).Methods("GET")
	router.HandleFunc("/{id}", PostController.Update).Methods("PATCH")
	router.HandleFunc("/{id}", PostController.Delete).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":9000", router))
}
