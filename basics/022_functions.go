package main

import "fmt"

func main() {

	// 1)
	// func <name>(parameters list) returnType {
	// code block
	// return value
	// }

	// 3)
	// sum := add(1, 2)
	// --OR--
	// fmt.Println(add(2, 3))

	// 4)
	// func() {
	// 	fmt.Println("Hello Anonymous Function")
	// }() // immediately invoked function (IFFI)

	// 5)
	// greet := func() {
	// 	fmt.Println("Hello Anonymous Function")
	// }

	// greet()

	// 6)
	// operation := add

	// result := operation(3, 5)

	// fmt.Println(result)

	// 8)
	// Passing a function as an argument
	result := applyOperation(5, 3, add)
	fmt.Println("5 + 3 =", result)

	// 10)
	// Returning and using a function
	multiplyBy2 := createMultiplier(2)
	fmt.Println("6 * 2 = ", multiplyBy2(6))

}

// 2)
func add(a, b int) int {
	return a + b
}

// 7)
// Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// 9)
// Function that returns a function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}
