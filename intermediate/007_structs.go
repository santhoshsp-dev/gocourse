package intermediate

import "fmt"

// Step: 1)
// type Person struct {
// 	firstName string
// 	lastName  string
// 	age       int
// }

// Step: 11)
// type Person struct {
// 	firstName string
// 	lastName  string
// 	age       int
// 	address   Address
// }

// Step: 16)
type Person struct {
	firstName     string
	lastName      string
	age           int
	address       Address
	PhoneHomeCell // embeded anonymous field
}

// Step: 15)
type PhoneHomeCell struct {
	home string
	cell string
}

// Step: 10)
type Address struct {
	city    string
	country string
}

// Step: 6)
func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

// Step: 8)
func (p *Person) incrementAgeByOne() {
	p.age++
}

func main() {

	// Step: 2)
	// p1 := Person{
	// 	firstName: "John",
	// 	lastName:  "Doe",
	// 	age:       30,
	// }

	// Step: 3)
	p2 := Person{
		firstName: "Jane",
		age:       25,
	}

	// Step: 19)
	p3 := Person{
		firstName: "Jane",
		age:       25,
	}

	// Step: 12)
	// p1 := Person{
	// 	firstName: "John",
	// 	lastName:  "Doe",
	// 	age:       30,
	// 	address: Address{
	// 		city:    "London",
	// 		country: "U.K.",
	// 	},
	// }

	// Step: 17)
	p1 := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		address: Address{
			city:    "London",
			country: "U.K.",
		},
		PhoneHomeCell: PhoneHomeCell{
			home: "465456454",
			cell: "45456464544",
		},
	}

	// Step: 13)
	// p2.address.city = "New York"
	// p2.address.country = "USA"

	// Step: 4)
	fmt.Println(p1.firstName)
	fmt.Println(p2.firstName)

	// Step: 7)
	fmt.Println(p1.fullName())

	// Step: 14)
	fmt.Println(p1.address.city)
	fmt.Println(p1.address)
	fmt.Println(p2.address.country)

	// Step: 18)
	fmt.Println(p1.cell)
	fmt.Println(p1.address.city)
	fmt.Println("Are p1 and p2 equal:", p1 == p2)

	// Step: 20)
	fmt.Println("Are p3 and p2 equal:", p3 == p2) // to run this code need to comment the Step: 13

	// Step: 5)
	// Anonymous Structs
	user := struct {
		username string
		email    string
	}{
		username: "user123",
		email:    "pseudoemail@example.org",
	}
	fmt.Println(user.username)

	// Step: 9)
	fmt.Println("Before increment", p1.age)
	p1.incrementAgeByOne()
	fmt.Println("After increment", p1.age)

}
