package main

import "fmt"

type employeeGoogle struct {
	firstName string
	lastName  string
	age       int
}

type employeeApple struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	const MAXRETRIES = 5
	var employeeID = 101
	fmt.Println("EmployeeID:", employeeID)
}
