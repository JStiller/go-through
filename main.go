package main

import "fmt"

// How to set constants
const a string = "hello world!"

func main() {
	fmt.Println(a)

	// How to set Variables
	var b [10]int
	fmt.Println(b)

	// Writing an for loop with by using an shorthand for n
	for n := 0; n < 10; n++ {
		b[n] = n
		switch {
		case b[n]%2 == 0:
			fmt.Println(b[n], "is even")
		default:
			fmt.Println(b[n], "is odd")
		}
	}

	fmt.Println(b)
}
