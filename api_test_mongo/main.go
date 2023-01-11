package main

import (
	"api_test_crud_mongo/handles"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customers/{id}", handles.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", handles.InsertCustomer).Methods("POST")

	http.ListenAndServe(":8080", router)
}
