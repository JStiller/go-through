package main

import "fmt"

// How to set constants
const a string = "hello world!"

func main() {
	fmt.Println(a)

	// How to set Variables
	var b int = 1
	fmt.Println(b)

	var c = true
	fmt.Println(c)

	var d int
	fmt.Println(d)

	// Using shorthand for declaring and initializing a variable
	e := "hello world!"
	fmt.Println(e)

	f := 1
	fmt.Println(f)

	g := true
	fmt.Println(g)
}
