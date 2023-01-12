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

	customer.Active = true
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

func ReplaceCustomer(w http.ResponseWriter, r *http.Request) {
	values := mux.Vars(r)
	customerId := values["id"]

	decoder := json.NewDecoder(r.Body)
	var customer entities.Customer

	err := decoder.Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	id, err := primitive.ObjectIDFromHex(customerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	customer.ID = id
	customer.Active = true
	customerReplaced, err := repositories.RaplaceCustomer(customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	if customerReplaced.ID == primitive.NilObjectID {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customerReplaced)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	values := mux.Vars(r)
	customerId := values["id"]

	err := repositories.DeleteCustomer(customerId)

	if err != nil && err.Error() == "Document not exist" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
