package main

import (
	"fmt"

	"github.com/vishwa-vijay/go/raja"
)

func main() {
	fmt.Println("Hello there!")
	io := raja.Io{}
	array := io.ReadSimpleArray("Please enter an array")
	for i, k in array
}
