package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

var persons = make([]Person, 0)

type Person struct {
	Name       string `json:"name"`
	MotherName string `json:"motherName"`
	Document   string `document`
}

func SavePerson(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var person Person
	err := decoder.Decode(&person)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}

	if person.Document == "" || person.MotherName == "" || person.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("The fields document, motherName and name are required.")
		return
	}

	persons = append(persons, person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(person)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {

	if len(persons) == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	values := mux.Vars(r)
	document := values["document"]

	hasPerson := false
	for index, person := range persons {
		if person.Document == document {
			hasPerson = true
			persons = append(persons[:index], persons[index+1:]...)
			break
		}
	}

	if !hasPerson {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
