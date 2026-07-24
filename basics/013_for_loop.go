package main

import "fmt"

func main() {
	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(i)
	// }

	// numbers := []int{1, 2, 3, 4, 5, 6}
	// for i, value := range numbers {
	// 	fmt.Printf("Index: %d | Value: %d\n", i, value)
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("Odd Number:", i)
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= rows; i++ {
	// 	for j := 1; j <= rows-i; j++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for k := 1; k <= 2*i-1; k++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

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

// for i := range 10 {

// 	fmt.Println(i)
// }
