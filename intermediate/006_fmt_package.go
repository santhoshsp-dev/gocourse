package main

import "fmt"

func main() {

	// Step: 1)
	// Printing Functions
	// fmt.Print("Hello ")
	// fmt.Print("World!")
	// fmt.Print(12, 456)

	// Step: 2)
	// fmt.Println("Hello ")
	// fmt.Println("World!")
	// fmt.Println(12, 456)

	// Step: 3)
	// name := "John"
	// age := 25
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	// fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// Step: 4)
	// Formatting Functions
	// s := fmt.Sprint("Hello", "World!", 123, 456)
	// fmt.Print(s)

	// Step: 5)
	// s = fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Print(s)
	// fmt.Print(s)

	// Step: 6)
	// sf := fmt.Sprintf("Name: %s, Age %d", name, age)
	// fmt.Println(sf)
	// fmt.Println(sf)

	// Step: 7)
	// Scanning Functions
	// var name string
	// var age int

	// Step: 8)
	// fmt.Print("Enter your name and age:")
	// // fmt.Scan(&name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Step: 9)
	// fmt.Print("Enter your name and age:")
	// // fmt.Scanln(&name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Step: 10)
	// fmt.Print("Enter your name and age:")
	// // fmt.Scan(&name, &age)
	// // fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Step: 12)
	// Error Formatting Functions
	err := checkAge(19)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

// Step: 11)
func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}
