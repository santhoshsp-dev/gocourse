package advanced
package main

// func main() {
// Step: 1)
// ch1 := make(chan int)
// ch2 := make(chan int)

// msg1 := <-ch1
// fmt.Println("Received from ch1:", msg1)
// msg2 := <-ch2
// fmt.Println("Received from ch1:", msg2)

// Step: 2)
// ch1 := make(chan int)
// ch2 := make(chan int)

// select {
// case msg := <-ch1:
// 	fmt.Println("Received from ch1:", msg)
// case msg := <-ch2:
// 	fmt.Println("Received from ch2:", msg)
// default:
// 	fmt.Println("No channels ready...")
// }
// fmt.Println("End of program")

// Step: 3)
// ch1 := make(chan int)
// ch2 := make(chan int)

// go func() {
// 	time.Sleep(time.Second)
// 	ch1 <- 1
// }()
// go func() {
// 	time.Sleep(time.Second)
// 	ch2 <- 2
// }()

// time.Sleep(2 * time.Second)
// select {
// case msg := <-ch1:
// 	fmt.Println("Received from ch1:", msg)
// case msg := <-ch2:
// 	fmt.Println("Received from ch2:", msg)
// default:
// 	fmt.Println("No channels ready...")
// }
// fmt.Println("End of program")

// Step: 4)
// 	ch1 := make(chan int)
// 	ch2 := make(chan int)

// 	go func() {
// 		time.Sleep(time.Second)
// 		ch1 <- 1
// 	}()
// 	go func() {
// 		time.Sleep(time.Second)
// 		ch2 <- 2
// 	}()

// 	time.Sleep(2 * time.Second)
// 	for range 2 {
// 		select {
// 		case msg := <-ch1:
// 			fmt.Println("Received from ch1:", msg)
// 		case msg := <-ch2:
// 			fmt.Println("Received from ch2:", msg)
// 		default:
// 			fmt.Println("No channels ready...")
// 		}
// 	}
// 	fmt.Println("End of program")
// }

// Step: 5)
// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		close(ch)
// 	}()

// 	select {
// 	case msg := <-ch:
// 		fmt.Println("Received", msg)
// 	case <-time.After(3 * time.Second):
// 		fmt.Println("Timeout")
// 	}
// }
