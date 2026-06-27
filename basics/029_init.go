package basics

import "fmt"

// 1)
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

	// 2)
	fmt.Println("Inside the main function")

}
