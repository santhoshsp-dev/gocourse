package main

import "fmt"

// Step: 12)
type Shape struct {
	Rectangle // embedd Rectangle in shape struct
}

// Step: 1)
type Rectangle struct {
	length float64
	width  float64
}

// Step: 2)
// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.length * r.width
}

// Step: 5)
// Method with pointer receiver
func (r *Rectangle) Scale(factor float64) {
	r.length *= factor // r.length = r.length * factor
	r.width *= factor
}

func main() {

	// Step: 3)
	rect := Rectangle{
		length: 10,
		width:  9,
	}

	// Step: 4)
	area := rect.Area()
	fmt.Println("Area of rectangle with width 9 and length 10 is", area)

	// Step: 6)
	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of rectangle with a factor of 2 is", area)

	// Step: 9)
	num := MyInt(-5)
	num1 := MyInt(9)
	fmt.Println(num.IsPositive())
	fmt.Println(num1.IsPositive())

	// Step: 11)
	fmt.Println(num.welcomeMessage())

	// Step: 13)
	s := Shape{Rectangle: Rectangle{length: 10, width: 9}}
	fmt.Println(s.Area())           // Area method is for rectangle. if we were to call area it would be [s.rectangle.area]. But now we are embedding a struct inside another struct, the method associated with the embedded struct will be promoted to the outer struct directly. so we can access Area directly like [s.Area]
	fmt.Println(s.Rectangle.Area()) // both are correct and give same output
}

// Step: 7)
type MyInt int

// Step: 8)
// Method on a user-defined type
func (m MyInt) IsPositive() bool {
	return m > 0
}

// Step: 10)
func (MyInt) welcomeMessage() string {
	return "Welcome to MyInt Type"
}
