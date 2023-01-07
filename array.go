package main

import "fmt"

func main() {
	// Here you define the length of the array
	var x [5]int

	x[2] = 6
	fmt.Printf("%d", x[2])

	// Array literal
	// var y [5]int = [5]{1, 2, 3, 4, 5}

	// Here you don't define the length of the array, it gets the length of the array from the number of elements you give it
	// z := [...] int {1, 2, 3, 4}
}
