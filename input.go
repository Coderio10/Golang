package main

import "fmt"

// Request input from user
func main() {

	var appleNum int
	fmt.Printf("Number Of Apples?")
	// demands input
	fmt.Scanf("%d", &appleNum)
	// return the input
	fmt.Printf("%d", appleNum)
}
