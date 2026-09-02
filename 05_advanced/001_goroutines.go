package main

import (
	"fmt"
	"time"
)

func main() {
	var err error
	fmt.Println("main(): Before sayHello() | ", time.Now())
	go sayHello()
	fmt.Println("main(): After sayHello(), before (2 Sec Wait) | ", time.Now())

	go func() {
		err = doWork()
	}()

	go printNumbers()
	fmt.Println("printNumbers(): After printNumbers(), before printLetters() | ", time.Now())
	go printLetters()
	fmt.Println("printLetters(): After printLetters(), before (2 * time.Second) | ", time.Now())

	time.Sleep(2 * time.Second)
	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Work completed successfully")
	}

	// time.Sleep(2 * time.Second)
	fmt.Println("main(): After (2 * time.Second) | ", time.Now())
}

func sayHello() {
	fmt.Println("sayHello(): Before (1 * time.Second) | ", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("sayHello(): After (1 * time.Second) | ", time.Now())
}

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println("Number: ", i, " | ", time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println("Letter: ", string(letter), " | ", time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	time.Sleep(1 * time.Second)
	return fmt.Errorf("An error occured in doWork.")
}
