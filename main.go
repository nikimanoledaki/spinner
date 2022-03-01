package main

import (
	"net/http"

	"github.com/callisto13/spinner/pkg/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/start", handlers.Start)
	r.HandleFunc("/stop", handlers.Stop)

	_ = http.ListenAndServe(":8080", r)
}
