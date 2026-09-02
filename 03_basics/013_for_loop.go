package main

import "fmt"

func main() {
	// Step: 1)
	// Simple iteration over a range
	// for initialization;condition;post{code block to be executed repeatedly}
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// Step: 2)
	// iterate over collection
	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for index, value := range numbers {
	// 	fmt.Printf("Index: %d, Value:%d\n", index, value) // %d specific to numbers
	// // --OR--
	// 	fmt.Printf("Index: %v, Value:%v\n", index, value) // %v general value
	// }

	// Step: 3)
	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue // continue the loop but skip the rest of lines/statements
	// 	}
	// 	fmt.Println("Odd Number:", i)
	// 	if i == 5 {
	// 		break // break out of the loop
	// 	}
	// }

	// Step: 4)
	// ASTERISK LAYOUT
	// rows := 5

	// //Outer loop
	// for i := 1; i <= rows; i++ {
	// 	// inner loop for spaces before stars
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	// inner loop for stars
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println() // Move to the next line
	// }

	// Step: 5)
	// Go 1.22 update
	// for i := range 10 {
	// 	i++
	// 	fmt.Println(i)
	// }
	// fmt.Println("We have a lift off!")

	//------------------------------------
	// rows := 5
	// for i := 1; i <= rows; i++ {
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := rows - 1; i >= 1; i-- {
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// 	// if i == 1 {
	// 	// 	break
	// 	// }
	// }
	//----------------------------------

	// rows := 4
	// for i := 1; i <= rows; i++ {
	// 	fmt.Print(" ")
	// 	for k := 1; k <= i; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	// for i := rows - 1; i >= 1; i-- {
	// 	fmt.Print(" ")
	// 	for k := 1; k <= i; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }
	// ----------------------------

	// rows := 5
	// for i := 1; i <= rows; i++ {
	// 	fmt.Print(" ")
	// 	for k := 1; k <= i; k++ {
	// 		fmt.Print(k)
	// 	}
	// 	fmt.Println()
	// }

	// for i := rows - 1; i >= 1; i-- {
	// 	fmt.Print(" ")
	// 	for k := 1; k <= i; k++ {
	// 		fmt.Print(k)
	// 	}
	// 	fmt.Println()
	// }
	//---------------------------------
	rows := 4
	letters := "ABCDE"

	for i := rows; i >= 0; i-- {
		// fmt.Print(" ")
		for k := 0; k <= i; k++ {
			fmt.Print(string(letters[k]))
		}
		fmt.Println()
	}
}
