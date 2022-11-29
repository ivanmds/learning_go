package main

import "fmt"

func calcTwoNumbers(x, y int) int {
	return x + y
}

func sumAllNumbers(args ...int) int {

	sumTotal := 0
	totalItens := len(args)

	for i := 0; i < totalItens; i++ {
		sumTotal = sumTotal + args[i]
	}

	return sumTotal
}

func main() {

	fmt.Println("Hello Ivan")
	fmt.Println("The result from sum between 2 and 6 is", calcTwoNumbers(2, 6))

	fmt.Println("Total is", sumAllNumbers(1, 8, 6, 97, 8))
}
