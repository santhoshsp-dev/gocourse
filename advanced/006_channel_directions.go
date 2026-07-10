package advanced
package advanced

import "fmt"

// Step: 1)
// func main() {
// 	ch := make(chan int)

// 	go func(ch chan<- int) {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}(ch)

// 	for value := range ch {
// 		fmt.Println("Received: ", value)
// 	}
// }

// Step: 2)
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := range 5 {
// 			ch <- i
// 		}
// 		close(ch)
// 	}()

// 	// for value := range ch {
// 	// 	fmt.Println("Received: ", value)
// 	// }
// 	receiveData(ch)
// }

// func receiveData(ch <-chan int) {
// 	for value := range ch {
// 		fmt.Println("Received: ", value)
// 	}
// }

// Step: 3)
func main() {

	ch := make(chan int)
	producer(ch)
	consumer(ch)
}

// Send only channel
func producer(ch chan<- int) {
	go func() {
		for i := range 5 {
			ch <- i
		}
		close(ch)
	}()
}

// Receive only channel
func consumer(ch <-chan int) {
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}
