package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func main() {

// 	// Step: 1)
// 	// i := 1
// 	// for i <= 5 {
// 	// 	fmt.Println("Iteration: ", i)
// 	// 	i++
// 	// }

// 	// Step: 2)
// 	// for as while with break
// 	// sum := 0
// 	// for {
// 	// 	sum += 10
// 	// 	fmt.Println("Sum:", sum)
// 	// 	if sum >= 50 {
// 	// 		break
// 	// 	}
// 	// }

// 	// Step: 3)
// 	num := 1
// 	for num <= 10 {
// 		if num%2 == 0 {
// 			num++
// 			continue
// 		}
// 		fmt.Println("Odd Number:", num)
// 		num++ // ++ increment operator increases value by 1 and -- decrement operator, decreases value by 1
// 	}

// }

// Step: 4)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	// Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have chosen a number between 1 and 100")
	fmt.Println("Can you guess what it is?")

	var guess int
	for {
		fmt.Println("Enter your guess: ")
		fmt.Scanln(&guess)

		// Check if the guess if correct
		if guess == target {
			fmt.Println("Congratulations! You guessed the correct number!")
			break
		} else if guess < target {
			fmt.Println("Too low! Try guessing a higher number.")
		} else {
			fmt.Println("Too high! Try guessing a lower number.")
		}
	}
}
