package main

import "fmt"

func main() {
	// fruit := "apple"

	// switch fruit {
	// case "orange":
	// 	fmt.Println("Orange")
	// case "apple":
	// 	fmt.Println("Apple")
	// default:
	// 	fmt.Println("Not a fruit")
	// }

	// day := "monday"

	// switch day {
	// case "monday":
	// 	fmt.Println("Its Monday")
	// case "tuesday":
	// 	fmt.Println("Its tuesday")
	// case "wednesday":
	// 	fmt.Println("Its wednesday")
	// case "thursday":
	// 	fmt.Println("Its thursday")
	// case "friday":
	// 	fmt.Println("Its friday")
	// case "saturday":
	// 	fmt.Println("Its saturday")
	// case "sunday":
	// 	fmt.Println("Its sunday")
	// default:
	// 	fmt.Println("Not a day")
	// }

	// mark := 82

	// switch {
	// case mark >= 90:
	// 	fmt.Println("Grade A")
	// case mark >= 80 && mark < 90:
	// 	fmt.Println("Grade B")
	// case mark > 70:
	// 	fmt.Println("Grade C")
	// case mark > 60:
	// 	fmt.Println("Grade D")
	// default:
	// 	fmt.Println("Failed")
	// }

	// num := 2
	// switch {
	// case num > 1:
	// 	fmt.Println("Grater than 1")
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not two")
	// }

	checkType(10)
	checkType("hello")
	checkType(34.4)
	checkType(true)
}

func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	case float64:
		fmt.Println("Float64")
	default:
		fmt.Println("Unknown type")
	}
}
