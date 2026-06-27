package main

import "fmt"

func main() {

	// 2)
	process()
	fmt.Println("Returned from process")

}

// 1)
// func process() {
// 	fmt.Println("Start Process")
// 	panic("Something went wrong!")
// 	fmt.Println("End Process")
// }

// 3)
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
