package basics

import (
	"fmt"
	"os"
)

func main() {
	// 4)
	defer fmt.Println("Deferred statement")

	// 1)
	fmt.Println("Starting the main function")

	// 2)
	// Exit with status code of 1
	os.Exit(1)

	// 3)
	// This will never be executed
	fmt.Println("End of main function")

}
