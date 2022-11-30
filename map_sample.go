package main

import "fmt"

func main() {
	fmt.Println("Some samples using go map")

	var nationCapitals map[string]string = make(map[string]string)
	nationCapitals["Brazil"] = "Brasília"
	nationCapitals["Japan"] = "Tokyo"
	nationCapitals["Taiwan"] = "Taipei"
	nationCapitals["Canada"] = "Ottawa"

	for key, value := range nationCapitals {
		fmt.Println("The capital of", key, "is", value)
	}
}
