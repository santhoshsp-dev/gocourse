package main

import (
	"errors"
	"fmt"
)

func main() {
	// q, r := divide(10, 3)
	// fmt.Printf("Quotient: %d, Redminder: %d\n", q, r)
	result, err := compare(3, 3)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println(result)
	}

}

func divide(a, b int) (int, int) {
	quotient := a / b
	reminder := a % b

	return quotient, reminder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "b is greater than a", nil
	} else {
		return "", errors.New("Unable to compare which is greater")
	}
}
