package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	stringToReverse := "Hello, OTUS!"
	outputString := stringutil.Reverse(stringToReverse)
	fmt.Println(outputString)
}
