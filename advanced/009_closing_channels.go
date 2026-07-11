package advanced
package advanced

import "fmt"

// func main() {
// Step: 1)
// === Simple closing channel example
// ch := make(chan int)

// go func() {
// 	for i := range 5 {
// 		ch <- i
// 	}
// 	close(ch)
// }()

// for val := range ch {
// 	fmt.Println(val)
// }

// Step: 2)
// RECEIVING FROM A CLOSED CHANNEL
// ch := make(chan int)
// close(ch)

// val, ok := <-ch
// if !ok {
// 	fmt.Println("Channel is closed")
// 	return
// }
// fmt.Println(val)

// -- OR--

// Step: 3)
// ch := make(chan int)
// 	close(ch)

// 	val, ok := <-ch
// 	if !ok {
// 		fmt.Println("Channel is closed")
// 	}else{
// 		fmt.Println(val)
// 	}

// Step: 4)
// RANGE OVER CLOSED CHANNEL
// ch := make(chan int)

// go func() {
// 	for i := range 5 {
// 		ch <- i
// 	}
// 	close(ch)
// }()

// for val := range ch {
// 	fmt.Println(val)
// }

// Step: 5)
// CLose channel twise will show error
// ch := make(chan int)
// go func() {
// 	close(ch)
// 	close(ch)
// }()
// time.Sleep(time.Second)
// }

// Step: 6)

func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)
}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for val := range ch2 {
		fmt.Println(val)
	}
}
