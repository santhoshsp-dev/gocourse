package main

import (
	"fmt"
	"time"
)

func main() {

	// Step: 1)
	//variable := make(chan type) '<-' operator
	greeting := make(chan string)
	greetString := "Hello"

	// Step: 2)
	// greeting <- gretString
	// receiver := <-greeting
	// fmt.Println(receiver)

	// Step: 3)
	// go func() {
	// 	greeting <- gretString // blocking because it is continuously trying to receive values, it is ready to receive continuous flow of data.
	// }()
	// receiver := <-greeting
	// fmt.Println(receiver)
	// fmt.Println("End of program.")

	// Step: 4)
	// go func() {
	// 	greeting <- gretString
	greeting <- "world"
	// }()
	// receiver := <-greeting
	// fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)
	// fmt.Println("End of program.")

	// Step: 5)
	// this code wont work, so we need to give some time to run
	// go func() {
	// 	greeting <- gretString
	// }()

	go func() {
		receiver := <-greeting
		fmt.Println(receiver)
	}()
	// fmt.Println("End of program.")

	// Step: 6)
	// go func() {
	// 	greeting <- gretString
	// }()

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// }()
	time.Sleep(1 * time.Second)
	// fmt.Println("End of program.")

	// Step: 7)
	// go func() {
	// 	greeting <- gretString
	greeting <- "world"
	// }()

	go func() {
		receiver := <-greeting
		fmt.Println(receiver)
		receiver = <-greeting
		fmt.Println(receiver)
	}()
	// time.Sleep(1 * time.Second)
	// fmt.Println("End of program.")

	// Step: 8)
	go func() {
		greeting <- greetString // blocking because it is continuously trying to receive values, it is ready to receive continuous flow of data.
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of program.")

}
