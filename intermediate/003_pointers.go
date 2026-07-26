// package main

// import "fmt"

// func main() {
// 	var ptr *int
// 	fmt.Println(ptr)
// 	a := 10
// 	ptr = &a
// 	fmt.Println(a)
// 	fmt.Println(ptr)
// 	fmt.Println(*ptr)
// 	fmt.Println("-------------------")
// 	fmt.Println(a)
// 	modifyValue(ptr)
// 	fmt.Println(a)
// 	fmt.Println(ptr)
// 	fmt.Println(*ptr)
// }

//	func modifyValue(ptr *int) {
//		*ptr++
//	}
//
// -------------------------------------
package main

import "fmt"

func main() {

	// Step: 1)
	var ptr *int
	var a int = 10
	ptr = &a // referencing

	fmt.Println(a)
	fmt.Println(ptr)

	// Step: 2)
	// fmt.Println(*ptr) // dereferencing a pointer

	// Step: 3)
	// var ptr *int
	// var a int = 10
	// fmt.Println(a)
	// if ptr == nil {
	// 	fmt.Println("Pointer is nil")
	// }

	// Step: 5)
	modifyValue(ptr)
	fmt.Println(a)

}

// Step: 4)
func modifyValue(ptr *int) {
	*ptr++
}
