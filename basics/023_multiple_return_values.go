package main

import (
	"errors"
	"fmt"
)

func main() {

	// Step: 1)
	// func functionName(parameter1 type1, parameter2 type2,...) (returnType1, returnType2,...){
	//code block
	// return returvalue1, returnValue2,...
	// }

	// Step: 3)
	q, r := divide(10, 3)
	fmt.Printf("Quotient: %v. Remainder: %v\n", q, r)

	// Step: 5)
	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}

// Step: 2)
// func divide(a, b int) (int, int) {
// 	quotient := a / b
// 	remainder := a % b
// 	return quotient, remainder
// }

// Step: 6
func divide(a, b int) (quotient int, remainder int) {
	quotient = a / b
	remainder = a % b
	return
}

// Step: 4)
func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if b > a {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}
