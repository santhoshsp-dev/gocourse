package main

import "fmt"

func main() {
	// var numbers [5]int
	// fmt.Println(numbers)

	// numbers[4] = 20
	// fmt.Println(numbers)

	// numbers[len(numbers)-1] = 30
	// fmt.Println(numbers)

	// fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}
	// fmt.Println("Fruits:", fruits)

	// originalArray := [3]int{1, 2, 3}
	// copyArray := originalArray

	// copyArray[0] = 100

	// fmt.Println(originalArray)
	// fmt.Println(copyArray)

	// for i := 1; i < len(numbers); i++ {
	// 	fmt.Println(i, ":", numbers[i])
	// }

	// for index, value := range numbers {
	// 	fmt.Println(index, ":", value)
	// }

	// var matrix [3][3]int = [3][3]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// }
	// fmt.Println(matrix)

	originalArray := [3]int{1, 2, 3}
	var copiedArray *[3]int

	copiedArray = &originalArray
	copiedArray[0] = 100

	fmt.Println("Orginal Array: ", originalArray)
	fmt.Println("Copied Array: ", *copiedArray)

}
