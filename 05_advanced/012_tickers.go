package main

// Step: 1)
// func main() {
// 	ticker := time.NewTicker(time.Second)

// 	for tick := range ticker.C {
// 		fmt.Println("Tick at:", tick)
// 	}
// }

// Step: 2)
// func main() {
// 	ticker := time.NewTicker(time.Second)

// 	i := 0
// 	for range ticker.C {
// 		i++
// 		fmt.Println(i)
// 	}
// }

// Step: 3)
// func main() {
// 	ticker := time.NewTicker(2 * time.Second)

// 	i := 1
// 	for range ticker.C {
// 		i *= 2
// 		fmt.Println(i)
// 	}
// }

// Step: 4)
// func main() {
// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()

// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println(i)
// 	}

// // 	for tick := range ticker.C {
// // 		i *= 2
// // 		fmt.Println(tick)
// // 	}
// }

// Step: 5)
// ========= SCHEDULING LOGGING, PERIODIC TASKS, POLLING FOR UPDATES
// func periodicTask() {
// 	fmt.Println("Performing periodic task at:", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// Step: 6)
// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	stop := time.After(5 * time.Second)

// 	for {
// 		select {
// 		case tick := <-ticker.C:
// 			fmt.Println("Tick at:", tick)
// 		case <-stop:
// 			fmt.Println("Stopping ticker.")
// 			return
// 		}
// 	}
// }
