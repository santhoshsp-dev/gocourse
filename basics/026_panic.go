package main

import "fmt"

func main() {

	// 1)
	// panic(interface{})

	// 3)
	// Example of a valid input
	process(10)

	// 4)
	// Example of an invalid input
	process(-3)
}

// 2)
func process(input int) {
	if input < 0 {
		panic("input must be a non-negative number")
	}
	fmt.Println("Processing input:", input)
}

// 5)
func process(input int) {
	defer fmt.Println("Deferred 1")
	defer fmt.Println("Deferred 2")
	if input < 0 {
		fmt.Println("Before Panic")
		panic("input must be a non-negative number")
		// fmt.Println("After Panic") // cant place this code here, Warning: unreachable code [code cant reach when it hit panic]
		// defer fmt.Println("Deferred 3") // cant place this code here, Warning: unreachable code [code cant reach when it hit panic]
	}
	fmt.Println("Processing input:", input)
}
