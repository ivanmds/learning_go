package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var input *bufio.Reader

func getLine() {
	for {
		line, err := input.ReadString('\n')

		if err != nil {
			log.Fatal(err)
		}

		if strings.Contains(line, "close") {
			fmt.Println("The program is closing")
			break
		}

		valueReverted := stringReverter(line)
		fmt.Println(valueReverted)
	}
}

func stringReverter(value string) string {
	valueReverted := ""

	for _, v := range value {
		valueReverted = string(v) + valueReverted
	}

	return valueReverted
}

func main() {
	fmt.Println("Hello, write a text to reverter")
	input = bufio.NewReader(os.Stdin)
	getLine()
}
