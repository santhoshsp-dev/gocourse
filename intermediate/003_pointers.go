package main

import "fmt"

func main() {

	// 1)
	var ptr *int
	var a int = 10
	ptr = &a // referencing

	fmt.Println(a)
	fmt.Println(ptr)

	// 2)
	// fmt.Println(*ptr) // dereferencing a pointer

	// 3)
	// var ptr *int
	// var a int = 10
	// fmt.Println(a)
	// if ptr == nil {
	// 	fmt.Println("Pointer is nil")
	// }

	// 5)
	modifyValue(ptr)
	fmt.Println(a)

}

// 4)
func modifyValue(ptr *int) {
	*ptr++
}
