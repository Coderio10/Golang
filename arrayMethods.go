package main

import "fmt"

func main() {
	// ITERATING THROUGH ARRAYS using the range keyword
	// RANGE
	// Range has two values: index and Element at the index
	x := [...]int{1, 2, 3, 4}

	for i, v := range x {
		fmt.Printf("index: %d value: %d\n", i, v)
	}
	// SLICE
}
