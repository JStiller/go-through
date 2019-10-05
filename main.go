package main

import "fmt"

// How to set constants
const a string = "hello world!"

func main() {
	fmt.Println(a)

	// How to set Variables
	var b int = 1
	fmt.Println(b)

	// Using shorthand for declaring and initializing a variable
	i := 1

	// Writing a simple for loop
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// Writing a classic intial (j), condition (j > 0) and after (j--) for loop
	for j := 10; j > 0; j-- {
		fmt.Println(j)
	}

	// Wirting an for loop without any condition. This loop will loop repeatedly until you break out of the loop
	for {
		fmt.Println("break loop")
		break
	}

	// Writing an for loop with an continue. The counterpart of an break
	for odd := 0; odd <= 5; odd++ {
		if odd%2 == 0 {
			continue
		}

		fmt.Println(odd)
	}
}
