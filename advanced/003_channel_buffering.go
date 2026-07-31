package main

import (
	"fmt"
	"time"
)

// func main() {
// Step: 1)
// ch := make(chan int)
// ch <- 1

// receiver := <-ch
// fmt.Println(receiver)

// Step: 2)
// ch := make(chan int)

// go func() {
// 	ch <- 1
// }()

// receiver := <-ch
// fmt.Println(receiver)

// Step: 3)
// go func() {
// 	time.Sleep(2 * time.Second)
// 	fmt.Println("2 second Goroutine finished")
// }()

// go func() {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println("3 second Goroutine finished")
// }()

// receiver := <-ch
// fmt.Println(receiver)
// fmt.Println("End of program")

// Step: 4)
// go func() {
// 		ch <- 1
// 		time.Sleep(2 * time.Second)
// 		fmt.Println("2 second Goroutine finished")
// 	}()

// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		fmt.Println("3 second Goroutine finished")
// 	}()

// 	receiver := <-ch
// 	fmt.Println(receiver)
// 	fmt.Println("End of program")

// Step: 5)
// ch := make(chan int)
// 	go func() {
// 		time.Sleep(3 * time.Second)
// 		receiver := <-ch
// 		fmt.Println(receiver)
// 		fmt.Println("3 second Goroutine finished")
// 	}()
// 	ch <- 1
// fmt.Println("End of program")

// Step: 6)
// ch := make(chan int)
// go func() {
// 	time.Sleep(3 * time.Second)
// 	fmt.Println(<-ch)
// 	fmt.Println("3 second Goroutine finished")
// }()
// ch <- 1
// fmt.Println("End of program")
// }

// Step: 7)
// UNBUFFERED CHANNELS
func main() {

	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)

		fmt.Println(<-ch)
		fmt.Println("3 second Goroutine finished")
	}()
	ch <- 1
	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("2 second Goroutine finished")
	// }()
	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("3 second Goroutine finished")
	// }()
	// receiver := <-ch
	// fmt.Println(receiver)
	fmt.Println("End of program")
}
