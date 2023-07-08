package main

import (
	PostController "app/controllers"
	"net/http"
)

func main() {
	http.HandleFunc("/", PostController.Index)
	http.ListenAndServe(":9000", nil)
}
