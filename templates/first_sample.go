package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	Name string
}

func greetingHandler(w http.ResponseWriter, r *http.Request) {
	var name string
	name = r.URL.Query().Get("name")
	if name == "" {
		name = "Gopher"
	}

	person := Person{Name: name}
	renderTemplate(w, "./pages/greeting.html", person)

}

func renderTemplate(w http.ResponseWriter, templateFile string, person Person) {
	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template", err)
	}
	t.Execute(w, person)
}

func main() {
	http.HandleFunc("/greeting", greetingHandler)
	http.ListenAndServe(":8181", nil)
}
