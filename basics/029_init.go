package main

import "fmt"

// Step: 1)
func init() {
	fmt.Println("Initializing package1...")
}

func init() {
	fmt.Println("Initializing package2...")
}

func init() {
	fmt.Println("Initializing package3...")
}

func main() {

	// Step: 2)
	fmt.Println("Inside the main function")
}
