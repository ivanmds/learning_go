package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"test_mux.com/m/handlers"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/persons", handlers.GetPerson).Methods("GET")
	r.HandleFunc("/persons", handlers.SavePerson).Methods("POST")
	r.HandleFunc("/persons/{document}", handlers.DeletePerson).Methods("DELETE")

	http.ListenAndServe(":8088", r)
}
