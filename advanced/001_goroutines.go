package advanced

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished/ready to return any value
// Goroutines do not stop the program flow and are non blocking

func main() {
	// Step: 11)
	var err error

	// Step: 5)
	fmt.Println("Beginning program.")

	// Step: 2)
	// sayHello()

	// Step: 3)
	go sayHello()

	// Step: 6)
	fmt.Println("After sayHello function.")

	// Step: 14)
	go func() {
		err = doWork()
	}()

	// Step: 12)
	// err = go doWork() // This is not accepted

	// Step: 9)
	go printNumbers()
	go printLetters()

	// Step: 4)
	time.Sleep(2 * time.Second)

	// Step: 13)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}
}

// Step: 1)
func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

// Step: 7)
func printNumbers() {
	for i := 0; i < 5; i++ {
		// fmt.Println(i)
		fmt.Println("Number: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

// Step: 8)
func printLetters() {
	for _, letter := range "abcde" {
		// fmt.Println(string(letter))
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

// Step: 10)
func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)

	return fmt.Errorf("an error occured in doWork.")
}
