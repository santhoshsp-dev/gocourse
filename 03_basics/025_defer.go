// package main

// import "fmt"

// func main() {
// 	process(10)
// }

//	func process(i int) {
//		fmt.Println("Normal Top")
//		defer fmt.Println("Defer i", i)
//		defer fmt.Println("Defer 1")
//		defer fmt.Println("Defer 2")
//		defer fmt.Println("Defer 3")
//		defer fmt.Println("Defer Last i", i)
//		i++
//		fmt.Println("Normal Bottom")
//	}
package main

import "fmt"

func main() {

	// Step: 2)
	// process()
	process(10)

}

// Step: 1)
// func process() {
// 	defer fmt.Println("First deferred statement executed")
// 	defer fmt.Println("Second deferred statement executed")
// 	defer fmt.Println("Third deferred statement executed")
// 	fmt.Println("Normal execution statement")
// }

func process(i int) {
	defer fmt.Println("Deffered i value:", i)
	defer fmt.Println("First deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third deferred statement executed")
	i++
	fmt.Println("Normal execution statement")
	fmt.Println("Value of i:", i)
}
