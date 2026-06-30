package intermediate

import "fmt"

// Step: 1)
type person struct {
	name string
	age  int
}

// Step: 2)
// type Employee struct {
// 	person
// 	empId  string
// 	salary float64
// }

// Step: 8)
type Employee struct {
	employeeInfo person // Embedded struct Named field
	// person  // Anonymous field
	empId  string
	salary float64
}

// Step: 5)
// Introduce is a method that was associated with person struct. Introduce is a method person and not employee, but because person is embedded into employee, we can directly access inroduce from an instance of employee.
func (p person) introduce() {
	fmt.Printf("Hi, I'm %s and I'm %d years old.\n", p.name, p.age)
}

// Step: 7)
func (e Employee) introduce() { // the introduce method is now overridden by the employee method of the same name introduce. Now employees introduce method, overrides persons introduce method.
	// fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.name, e.empId, e.salary)
	fmt.Printf("Hi, I'm %s, employee ID: %s, and I earn %.2f.\n", e.employeeInfo.name, e.empId, e.salary)
}

func main() {

	// Step: 3)
	// emp := Employee{
	// 	person: person{name: "John", age: 30},
	// 	empId:        "E001", salary: 50000,
	// }

	// Step: 9)
	emp := Employee{
		employeeInfo: person{name: "John", age: 30},
		empId:        "E001", salary: 50000,
	}

	// Step: 4)
	// fmt.Println("Name:", emp.name) // Accessing the embedded struct field emp.person.name [field promotion]
	// fmt.Println("Age:", emp.age) // Same as above
	// fmt.Println("Emp ID:", emp.empId)
	// fmt.Println("Salary:", emp.salary)

	// Step: 10)
	fmt.Println("Name:", emp.employeeInfo.name)
	fmt.Println("Age:", emp.employeeInfo.age) // Same as above
	fmt.Println("Emp ID:", emp.empId)
	fmt.Println("Salary:", emp.salary)

	// Step: 6)
	emp.introduce()

}
