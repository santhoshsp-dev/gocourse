package main

import "fmt"

func main() {

	// Step: 2)
	process()
	fmt.Println("Returned from process")

}

// Step: 1)
// func process() {
// 	fmt.Println("Start Process")
// 	panic("Something went wrong!")
// 	fmt.Println("End Process")
// }

// Step: 3)
func process() {
	defer func() {
		// if r := recover(); r != nil {
		r := recover()
		if r != nil {
			fmt.Println("Recovered:", r)
		}
	}()

	fmt.Println("Start Process")
	panic("Something went wrong!")
	fmt.Println("End Process")
}
