package main

import "fmt"

func main() {
	// Step: 1)
	// if condition {
	// 	block of code
	// }

	// age := 25
	// if age >= 18 {
	// 	fmt.Println("You are an adult.")
	// }

	// Step: 2)
	// if condition {
	// block of code
	// } else if {
	// block of code
	// } else {
	// block of code
	// }

	// temperature := 25
	// if temperature >= 30 {
	// 	fmt.Println("It's hot outside.")
	// } else {
	// 	fmt.Println("It's cool outside.")
	// }

	// Step: 3)
	score := 95
	// if score >= 90 {
	// 	fmt.Println("Grade A")
	// } else if score >= 80 {
	// 	fmt.Println("Grade B")
	// } else if score >= 70 {
	// 	fmt.Println("Grade C")
	// } else {
	// 	fmt.Println("Grade D")
	// }

	// Step: 4)
	if score >= 90 {
		fmt.Println("Grade A")
	} else if score >= 80 {
		fmt.Println("Grade B")
	} else if score >= 70 {
		fmt.Println("Grade C")
	} else {
		fmt.Println("Grade D")
	}
	// anything here in this line will be executed after one of the condition is met

	// Step: 5)
	// nested if
	// if condition1 {
	// Code block1
	// if condition2{
	// code block 2
	// }
	// }

	// num := 18
	// if num%2 == 0 {
	// 	if num%3 == 0 {
	// 		fmt.Println("Number is divisible by both 2 and 3.")
	// 	} else {
	// 		fmt.Println("Number is divisible by 2 but not 3.")
	// 	}
	// } else {
	// 	fmt.Println("Number is not divisible by 2.")
	// }

	// Step: 6)
	// Logical operators
	// || OR
	// && AND

	if 10%2 == 0 || 6%2 == 0 {
		fmt.Println("Either 10 or 6 are even.")
	}

	// Step: 7)
	if 10%2 == 0 && 6%2 == 0 {
		fmt.Println("Both 10 and 6 are even.")
	}
}
