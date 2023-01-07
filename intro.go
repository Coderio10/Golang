package main

import "fmt"

// You can define a variable outside the function main
// don't do these
// 	var d string = "cat"
// 	var e string = "food"
// 	var f int = 34
// // do these instead 
// 	var {
// 		g int = 34
// 		h string = "food"
// 		j string = "cat"
// 	}
// 	var {
// 		counter int = 0
// 	}
// 	fmt.Printf(d)
// 	fmt.Printf(e)
// 	fmt.Printf(f)
// 	fmt.Printf(g)
// 	fmt.Println(h)
// 	fmt.Println(j)

func main(){
	fmt.Printf("Hello world")
// Variables
// 1st and 2nd method, you get to define the type of variable. The 3rd method, computer detects the type of variable 
	// 1st method
		var a int
		a = 27 
		fmt.Println(a)
	// 2nd method
		// integer
			var b int = 42
			fmt.Println(b)
		// float
			var b1 float32 = 25.5
			fmt.Printf("%v, %T\n", b1, b1) 
	// 3rd Method
		i := 42.
		fmt.Println(i)
		// detect variable
		fmt.Printf("%v, %T", i, i)

	// Scan: its used to request input from user  
}
