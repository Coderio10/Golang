package main

import "fmt"

func main() {
	// Defining Variable
	var deciNum float64
	// printing "Enter a decimal number"
	fmt.Printf("Enter a decimal number")
	// Demanding Input (Float)
	fmt.Scanf("%f", &deciNum)
	// Converting Input to integer and it prints it out
	fmt.Printf("The Truncated value of %f, is %d\n", deciNum, int(deciNum))
}
