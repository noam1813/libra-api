package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", echoParam)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func echoParam(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	param := r.URL.RawQuery
	fmt.Fprintf(w, param)
}
