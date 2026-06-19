package main

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	const MAXRETRIES = 5

	var employeeID = 1001
	fmt.Println("EmployeeID: ",employeeID)
}