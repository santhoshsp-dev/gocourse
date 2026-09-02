package main

// Step: 1)
// func main() {
// fmt.Println("Starting app.")
// timer := time.NewTimer(2 * time.Second) // time.NewTimer nonblocking in nature but time.Sleep is blocking in nature
// fmt.Println("Waiting for timer.C")
// <-timer.C // Blocking in nature
// fmt.Println("Timer Expired")
// }

// Step: 2)
// func main(){
// fmt.Println("Starting app.")
// timer := time.NewTimer(2 * time.Second)
// fmt.Println("Waiting for timer.C")
// stopped := timer.Stop()
// if stopped {
// 	fmt.Println("Timer stopped")
// }
// time.Stop will show an error, so we need to comment 2 lines of codes below
// <-timer.C // Blocking in nature
// fmt.Println("Timer Expired")
// }

// Step: 3)
// // ======== BASIC TIMER USE
// func main() {
// fmt.Println("Starting app.")
// timer := time.NewTimer(2 * time.Second)
// fmt.Println("Waiting for timer.C")
// stopped := timer.Stop()
// if stopped {
// 	fmt.Println("Timer stopped")
// }
// fmt.Println("Timer Reset")
// timer.Reset(time.Second)
// <-timer.C // Blocking in nature
// fmt.Println("Timer Expired")
// }

// Step: 4)
// // ============= TIMEOUT
// func logRunningOperation() {
// 	for i := range 20 {
// 		fmt.Println(i)
// 		time.Sleep(time.Second)
// 	}
// }

// func main() {
// 	timeout := time.After(2 * time.Second)
// 	done := make(chan bool)

// 	go func() {
// 		logRunningOperation()
// 		done <- true
// 	}()

// 	select {
// 	case <-timeout:
// 		fmt.Println("Operation timed out")
// 	case <-done:
// 		fmt.Println("Operation completed")
// 	}
// }

// Step: 5)
// // =========== SCHEDULING DELAYED OPERATIONS
// func main() {
// 	timer := time.NewTimer(2 * time.Second)

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation executed")
// 	}()
// }

// Step: 6)
// func main() {
// 	timer := time.NewTimer(2 * time.Second) // non blocking timer starts

// 	go func() {
// 		<-timer.C
// 		fmt.Println("Delayed operation executed")
// 	}()
// 	fmt.Println("Waiting...")
// 	time.Sleep(3 * time.Second) // blocking timer starts
// 	fmt.Println("End of the program")
// }

// Step: 7)
// func main() {
// 	timer1 := time.NewTimer(1 * time.Second)
// 	timer2 := time.NewTimer(2 * time.Second)

// 	select {
// 	case <-timer1.C:
// 		fmt.Println("Timer 1 expired")
// 	case <-timer2.C:
// 		fmt.Println("Timer 2 expired")
// 	}
// }
