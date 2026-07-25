package main

import (
	"fmt"
	"os"
)

func main() {
	// Step: 4)
	defer fmt.Println("Deferred statement")

	// Step: 1)
	fmt.Println("Starting the main function")

	// Step: 2)
	// Exit with status code of 1
	os.Exit(1)

	// Step: 3)
	// This will never be executed
	fmt.Println("End of main function")
}
