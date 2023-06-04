package main

import (
	"fmt"

	sutil "golang.org/x/example/stringutil"
)

func main() {
	input := "Hello, OTUS!"
	input = sutil.Reverse(input)
	fmt.Println(input)
}
