package main

import "fmt"

func main() {

	// 2)
	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	// 4)
	fmt.Println(sumOfDigits(9))
	fmt.Println(sumOfDigits(12))
	fmt.Println(sumOfDigits(12345))

}

// 1)
func factorial(n int) int {
	// Base case: factorial of 0 is 1
	if n == 0 {
		return 1
	}
	// Recursive case: factorial of n is n * factorial(n - 1)
	return n * factorial(n-1)
	// n * (n - 1) * (n-2) * factorial (n-3)..... factorial(0)

}

// 3)
func sumOfDigits(n int) int {
	// Base case
	if n < 10 {
		return n
	}
	return n%10 + sumOfDigits(n/10)
}

// n = 11
// 11 % 10 = 1
// 11 / 10 = 1

// n = 12
// 12 % 10 = 2
// 12 / 10 = 1
