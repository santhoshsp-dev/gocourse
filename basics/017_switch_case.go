package main

import "fmt"

func main() {

	// Step: 1)
	// Switch statement in go is (switch case default) (fallthrough)
	// switch expression {
	// case value1:
	// Code to be executed if expression equals value1
	// fallthrough
	// case value2:
	// Code to be executed if expression equals value2
	// case value3:
	// Code to be executed if expression equals value3
	// default:
	// Code to be executed if expression does not match any value
	// }

	// Step: 2)
	// Switch statement in other languages (switch case break default)
	// switch expression {
	// case value1:
	// Code to be executed if expression equals value1
	// break;
	// case value2:
	// Code to be executed if expression equals value2
	// break;
	// case value3:
	// Code to be executed if expression equals value3
	// break;
	// default:
	// Code to be executed if expression does not match any value
	// break;
	// }

	// Step: 3)
	// fruit := "pineapple"

	// switch fruit {
	// case "apple":
	// 	fmt.Println("It's an apple.")
	// case "banana":
	// 	fmt.Println("It's a banana.")
	// default:
	// 	fmt.Println("Unknown Fruit!")
	// }

	// Step: 4)
	// Multiple Conditions
	// day := "Monday"

	// switch day {
	// case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
	// 	fmt.Println("It's a weekday.")
	// case "Sunday":
	// 	fmt.Println("It's a weekend.")
	// default:
	// 	fmt.Println("Invalid day.")
	// }

	// Step: 5)
	// number := 15

	// switch {
	// case number < 10:
	// 	fmt.Println("Number is less than 10")
	// case number >= 10 && number < 20:
	// 	fmt.Println("Number is between 10 and 19")
	// default:
	// 	fmt.Println("Number is 20 or more")
	// }

	// Step: 6)
	// num := 2

	// switch {
	// case num > 1:
	// 	fmt.Println("Greater than 1")
	// 	fallthrough
	// case num == 2:
	// 	fmt.Println("Number is 2")
	// default:
	// 	fmt.Println("Not Two")
	// }

	// Step: 8)
	checkType(10)
	checkType(3.14)
	checkType("Hello")
	checkType(true)

}

// Step: 7)
func checkType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("It's an integer")
	case int32:
		fmt.Println("It's an integer")
	case float64:
		fmt.Println("It's float")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Unknown Type")
	}
}
