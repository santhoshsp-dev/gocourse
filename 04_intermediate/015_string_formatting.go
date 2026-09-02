package main

import "fmt"

func main() {

	// Step: 1)
	num := 424
	fmt.Printf("%05d\n", num)

	// Step: 2)
	message := "Hello"
	fmt.Printf("|%10s|\n", message)
	fmt.Printf("|%-10s|\n", message)

	// Step: 3)
	message1 := "Hello \nWorld!"
	message2 := `Hello \nWorld!`

	fmt.Println(message1)
	fmt.Println(message2)

	// Step: 4)
	// sqlQuery := `SELECT * FROM users WHERE age > 30`

}
