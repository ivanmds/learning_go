package main

import "fmt"

func main() {
	fmt.Println("Slice sample")
	var mySlice = make([]int, 5)
	populateIntegerSlice(mySlice)
	fmt.Printf("My slice %v. With len = %d and cap = %d\n", mySlice, len(mySlice), cap(mySlice))

	fmt.Println("Add new item in slice(dynamic array)")
	mySlice = append(mySlice, 55)
	fmt.Printf("My slice with new item %v. With len = %d and cap = %d\n", mySlice, len(mySlice), cap(mySlice))
}

func populateIntegerSlice(slice []int) {
	slice[0] = 5
	slice[1] = 9
	slice[2] = 8
	slice[3] = 6
	slice[4] = 10
}
