package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello there")
}

func customGreeting(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "buddy"
	}

	greeting := "Hello " + name + ", how are you doing?"

	fmt.Fprintf(w, greeting)

}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/greeting", customGreeting)
	http.ListenAndServe(":8181", nil)
}
