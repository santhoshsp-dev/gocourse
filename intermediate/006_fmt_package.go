package intermediate

import "fmt"

func main() {

	// 1)
	// Printing Functions
	// fmt.Print("Hello ")
	// fmt.Print("World!")
	// fmt.Print(12, 456)

	// 2)
	// fmt.Println("Hello ")
	// fmt.Println("World!")
	// fmt.Println(12, 456)

	// 3)
	// name := "John"
	// age := 25
	// fmt.Printf("Name: %s, Age: %d\n", name, age)
	// fmt.Printf("Binary: %b, Hex: %X\n", age, age)

	// 4)
	// Formatting Functions
	// s := fmt.Sprint("Hello", "World!", 123, 456)
	// fmt.Print(s)

	// 5)
	// s = fmt.Sprintln("Hello", "World!", 123, 456)
	// fmt.Print(s)
	// fmt.Print(s)

	// 6)
	// sf := fmt.Sprintf("Name: %s, Age %d", name, age)
	// fmt.Println(sf)
	// fmt.Println(sf)

	// 7)
	// Scanning Functions
	// var name string
	// var age int

	// 8)
	// fmt.Print("Enter your name and age:")
	// // fmt.Scan(&name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// 9)
	// fmt.Print("Enter your name and age:")
	// // fmt.Scanln(&name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// 10)
	// fmt.Print("Enter your name and age:")
	// // fmt.Scan(&name, &age)
	// // fmt.Scanln(&name, &age)
	// fmt.Scanf("%s %d", &name, &age)
	// fmt.Printf("Name: %s, Age: %d\n", name, age)

	// 12)
	// Error Formatting Functions
	err := checkAge(19)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}

// 11)
func checkAge(age int) error {
	if age < 18 {
		return fmt.Errorf("Age %d is too young to drive.", age)
	}
	return nil
}
