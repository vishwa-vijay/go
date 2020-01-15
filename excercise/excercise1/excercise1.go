package main

import (
	"fmt"

	"gitlab.com/vishwavijay.com/smd/raja"
)

func main() {
	io := raja.NewIo()
	array, rows, cols := io.ReadArray("Enter 2 diamention array.")
	println("Array is ____")
	raja.PrintArray(array)
	marker := raja.CreatEmptyTwoDiamentionalArray(rows, cols)
	raja.PrintArray(marker)

	fmt.Println("Done")
	fmt.Println(array)
}
