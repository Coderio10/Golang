package main

import "fmt"

func main() {
	// loops
	// first method
	for i := 0; i < 10; i++ {
		fmt.Printf("Hi\n")
	}
	// second method
	j := 0
	for j < 10 {
		fmt.Printf("Code Rio\n")
		j++
	}
	// LOOP METHODS:
	// 1. Break: it breaks the loop if defined condition is met
	k := 0
	for k < 10 {
		if k == 5 {
			break
		}
		k++
		fmt.Printf("Hey\n")
	}
	// 2. coninue: it starts the loop if defined condition is met
	for l := 0; l < 10; l++ {
		if l == 5 {
			continue
		}
		fmt.Printf("NITDA\n")
	}
}
