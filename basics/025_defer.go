package basics

import "fmt"

func main() {

	// 2)
	// process()
	process(10)

}

// 1)
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
