package main

import "fmt"

func main(){
	// Boolean
		var t bool = true 
		fmt.Printf("%v, %T\n", t, t)
		
		n := 1 == 1
		m := 1 == 3 
		fmt.Printf("%v, %T\n", n, n)
		fmt.Printf("%v, %T\n", m, m)

	// Numberic
		//Integers
			// You can't work with two number of different types unless you convert one to the other
			var l int = 10
			var k int8 = 5
			fmt.Println(l + int(k)) 
			a := 10 // 1010
			b := 3 	// 0011
			fmt.Println(a + b)
			fmt.Println(a - b)
			fmt.Println(a * b)
			// it will divide and convert the result into an integer
			fmt.Println(a / b)
			// % means modulo (remainder)
			fmt.Println(a % b)

			// Bit Operators (AND, OR, XOR, XAND ). It calcultes in binary and gives result in decimal
			// AND
			fmt.Println(a & b) // 0010
			// OR
			fmt.Println(a | b) // 1011
			// XOR
			fmt.Println(a ^ b) // 1001
			// XAND
			fmt.Println(a &^ b) // 0100
			
			// Bit Shifting
			w := 8 // 2^3
			fmt.Println(w << 3) // 2^3 * 2^3 = 2*6
			fmt.Println(w >> 3) // 2^3 / 2^3 = 2^0 = 1
		// Floating Point	
			g := 3.14
			g = 13.7e72
			g = 2.1E14
			fmt.Printf("%v, %T", g, g)
			// Arithemetics with floats
				o := 88.5
				u := 95.5
				fmt.Println(o + u)
				fmt.Println(o - u)
				fmt.Println(o * u)
		//  Complex Numbers
		// we have: complex64 and complex128
			var p complex128 = 1 + 2i
			fmt.Printf("%v, %T\n", p, p)
}