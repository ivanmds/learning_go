package main

import (
	"fmt"

	"github.com/ivanmds/learning_go/tree/main/types_samples/simpleshape"
)

func main() {
	fmt.Println("test samples types")
	rectangle := simpleshape.NewRectangle(5, 9)
	triangle := simpleshape.NewTriangle(6, 2)

	fmt.Println("Area from rectangle", rectangle.Area())
	fmt.Println("Area from triangle", triangle.Area())
}
