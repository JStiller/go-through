package main

import "fmt"

// How to set constants
const a string = "hello world!"

func main() {
	fmt.Println(a)

	// How to set Variables
	var b int = 1
	fmt.Println(b)

	// Writing an for loop with by using an shorthand for n
	for n := 0; n <= 5; n++ {
		switch {
		case n%2 == 0:
			fmt.Println(n, "is even")
		default:
			fmt.Println(n, "is odd")
		}
	}
}
