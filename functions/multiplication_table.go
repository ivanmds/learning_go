package main

import "fmt"

func multiplyTwoNumber(x, y int) int {
	return x * y
}

func main() {
	fmt.Println("Let's beging")

	for x := 1; x <= 100; x++ {
		for y := 1; y <= 10; y++ {
			fmt.Printf("The result from %d x %d is %d \n", x, y, multiplyTwoNumber(x, y))
		}
		fmt.Println("")
	}
}
