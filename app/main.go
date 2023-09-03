package main

import (
	PostController "app/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	router := mux.NewRouter()

	corsHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8000"},
		AllowedMethods: []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})
	router.HandleFunc("/", PostController.Index).Methods("GET")
	router.HandleFunc("/", PostController.Create).Methods("POST")
	router.HandleFunc("/{id}", PostController.Show).Methods("GET")
	router.HandleFunc("/{id}", PostController.Update).Methods("PATCH")
	router.HandleFunc("/{id}", PostController.Delete).Methods("DELETE")

	// サーバー起動
	handler := corsHandler.Handler(router)
	log.Fatal(http.ListenAndServe(":9000", handler))
}
