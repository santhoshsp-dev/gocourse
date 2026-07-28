package main

// Step: 6)
import (
	"fmt"
	"math"
)

// Step: 1)
type geometry interface {
	area() float64 // its a blank function, unimplemented or undefined functions. undefined methods area and perim
	perim() float64
}

// Step: 3)
type circle struct { // structure
	radius float64
}

// Step: 5)
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Step: 8)
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Step: 9)
func (c circle) diameter() float64 {
	return 2 * c.radius
}

// Step: 2)
type rect struct {
	width, height float64
}

// Step: 4)
func (r rect) area() float64 { // rect struct implements a method that method we name it as area. area which is a function a method defined inside the geometry interface.
	return r.height * r.width
}

// Step: 7)
func (r rect) perim() float64 {
	return 2 * (r.height + r.width)
}

// Step: 13)
type rect1 struct {
	width, height float64
}

// Step: 14)
func (r rect1) area() float64 {
	return r.height * r.width
}

// Step: 17)
// func (r rect1) perim() float64 {
// 	return 2 * (r.height + r.width)
// }

// Step: 11)
/*
// Interface is useful to us in this way because rectangle and circle they have both The implemented the
functions area and param from the interface geometry.
Rect and circle can be used as geometry type measure function accepts geometry type of values as its
argument, but these are of type, rect and circle. These are structs with type, rect and circle.
All right. But as soon as they implemented area and param, they were able to use geometry interface as well.
They were able to associate themselves with geometry type and any function anywhere where geometry type
is accepted. Circle and rect would also be accepted over there, even though circle has an additional function.
So circle has an additional method associated with it which is diameter.
It doesn't matter. All circle needs to do. All a struct needs to do is implement all the methods defined in that interface.
And as soon as it does that, if it has any additional methods, it's okay. But these should also be implemented.
All the methods that are defined in an interface should be implemented by that struct to be able to
get associated with that interface type.
*/
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {

	// Step: 10)
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	// Step: 15)
	// r1 := rect1{width: 3, height: 4}

	// Step: 12)
	measure(r)
	measure(c)

	// Step: 16)
	// measure(r1)

	// Step: 19)
	myPrinter(1, "John", 45.9, true)

	// Step: 21)
	printType(9)
	printType("John")
	printType(false)

}

// Step: 20)
func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type: Int")
	case string:
		fmt.Println("Type: String")
	default:
		fmt.Println("Type: Unknown")
	}
}

// Step: 18)
func myPrinter(i ...interface{}) {
	for _, v := range i {
		fmt.Println(v)
	}
}
