package main

import (
	"api_test_crud_mongo/handles"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/customers", handles.FindCustomerPagination).Methods("GET")
	router.HandleFunc("/customers/{id}", handles.GetCustomer).Methods("GET")
	router.HandleFunc("/customers", handles.InsertCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", handles.ReplaceCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", handles.DeleteCustomer).Methods("DELETE")

	http.ListenAndServe(":8080", router)
}
