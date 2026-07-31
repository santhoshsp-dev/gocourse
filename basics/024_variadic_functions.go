package main

import "fmt"

func main() {

	// Step: 1)
	// ... Ellipsis
	// func functionName(param1 type1, param2 ...type2) returnType{
	// function body
	// }

	// Step: 3)
	// fmt.Println("Sum of 1, 2, 3:", sum(1, 2, 3))

	// Step: 5)
	// statement, total := sum("The sum of 1, 2, 3 is", 1, 2, 3, 4, 5, 6)
	// fmt.Println(statement, total)

	// Step: 7)
	sequence, total := sum(1, 20, 30, 40, 50, 60)
	fmt.Println("Sequence: ", sequence, "Total", total)
	sequence2, total2 := sum(2, 40, 36, 40, 50, 60)
	fmt.Println("Sequence: ", sequence2, "Total", total2)

	// Step: 8)
	numbers := []int{1, 2, 3, 4, 5, 9}

	sequence3, total3 := sum(3, numbers...)
	fmt.Println("Sequence: ", sequence3, "Total", total3)
}

// Step: 2)
// func sum(nums ...int) int {
// 	total := 0
// 	for _, v := range nums {
// 		total += v
// 	}
// 	return total
// }

// Step: 4)
// func sum(returnString string, nums ...int) (string, int) {
// 	total := 0
// 	for _, v := range nums {
// 		total += v
// 	}
// 	return returnString, total
// }

// Step: 6)
func sum(sequence int, nums ...int) (int, int) {
	total := 0
	for _, v := range nums {
		total += v
	}
	return sequence, total
}
