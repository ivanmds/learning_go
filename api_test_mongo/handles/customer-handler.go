package handles

import (
	"api_test_crud_mongo/entities"
	"api_test_crud_mongo/repositories"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetCustomer(w http.ResponseWriter, r *http.Request) {
	values := mux.Vars(r)
	customerId := values["id"]

	customer, err := repositories.GetCustomer(customerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	if customer.ID == primitive.NilObjectID {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

func InsertCustomer(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var customer entities.Customer

	err := decoder.Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	customerResult, err := repositories.InsertCustomer(customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customerResult)
}
