// 011 Timers

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

/*
A timer in go allows you to schedule an event to occur after a specified duration. It is useful for implementing timeouts, scheduling periodic tasks, and delaying operations. One of the key reasons for using timers is timeouts. We can implement timeout functionality to limit how long a particular operation should wait. Another reason would be to use delays in schedule operations to occur after a certain delay. We have been using Time.sleep, but we can use timers instead as well and we'll see how we can do that. Other than that, we can define periodic tasks using timers. These tasks execute recurringly at regular intervals. Now how do we create a timer? We create a timer using time package, and in the time package we use the new timer function. So let's create a new timer that will send the current time on its channel after a specified duration. So this is what timer does. It sends the current time after a time duration. So let's see that in action timer variable will hold time dot new timer. And let's give it a duration of two seconds. So two multiplied by time dot second. And now this timer sends the current time on its channel. So this timer is a struct. So the resulting value of this timer. Let's see that what it is. So new timer returns a struct of type timer. And this timer struct has a field C using which we can access the channel, which gives us time. So let's do that. We receive the time using receive channel, the receive operator of channels, and we access timer and then the see property timer dot. See that's it. And after that we print Fmt.println timer expired. All right. So now we already know that the receive channel will wait for a send channel to send value. And this is blocking in nature. So this is blocking in nature. And that's why next line will not execute until this channel receives a value. Let's see what happens. Go run timer. And here we go. Timer expired. Appeared before us only after timer Dorsey received a value. So let's introduce some more print statements. Starting app. Actually just will stop and then fmt.println waiting for timer dot. See? Let's see. So starting go and then waiting for timer dot see and then timer expired. So this timer dot see was blocking here. The see field of the timer type is a channel that receives the time when the timer expires and the timer expires after two seconds. Time.new timer has just one parameter that is time dot duration. And now in case if we want to stop the timer before it expires. How do we do that? Okay. And one more thing that we must have noticed by now is that time dot new timer does not block. Let's notice that again. Waiting for timer. See? So time dot new timer executed and then the next line executed. And we were blocked here. So time dot new timer is non-blocking in nature. All right. Now earlier before this lecture we were using time dot sleep to perform timeouts. Right. But time dot sleep is blocking in nature, right? So if I use time dot sleep and then time dot second. So it took one second for starting app, right? If I comment this and then run this we see starting up immediately. So time dot sleep is blocking in nature. Next lines of code will not be executed. However, time dot new timer is non-blocking in nature and we can use the timer dot see in a go routine as well so that this also becomes non-blocking. Go routine will go back to a different thread and the main execution will move on to the next line of code. So this is the advantage of using time dot new timer. And now let's come to a scenario where we want to stop the timer before it expires. So waiting for timer dot c and timer dot and we saw the option timer dot reset and timer dot stop. So we use timer dot stop and the timer will stop immediately. And what does timer dot stop return. So let's see what does it return. Timer dot stop. Doesn't take any argument. It returns bool. Stop. Prevents the timer from firing. It returns true if the call stops the timer. False if the timer has already expired or been stopped. So if you are stopping an already stopped timer. So obviously it won't be able to stop because there is no running timer. So let's now use this boolean value. So. Stopped. So stopped is going to store the Boolean value. That is that is being returned from timer dot stop. And if stopped is true we print a statement Fmt.println. Timer. Stopped. Timer. Save it. And let's run this first. Clear this, run this and timer stopped. And that's why we received an error. Because there was no sender. So let's remove this. And then. So timer stopped and instead of timer expired. I should comment this out as well because this print statement is stating timer expired because of the channel. All right. So this is how we stop a running timer before that timer expires. Now that we have experimented with timer dot stop, let's bring Timer Dorsey back. Let's also use timer dot reset. So timer dot reset. What it does is that it will restart a timer that has been expired or stopped. So we see here for a timer created with new timer reset should be invoked only on stopped or expired timers with drained channels. Perfect. So let's reset this timer for one second. So time dot second. Let's see if we get an error this time, because this gives us an error if the timer is stopped because there is no sender channel. And we have implemented a receiver channel here. So let's see the output. Perfect. Timer was stopped. But then we should use another print statement. So yeah. Timer reset. And then let's see. So timer stopped then timer reset. And after one second timer expired. So we have used timer, stop, timer, dot reset and timer dot see the channel that sends the expiry time perfect. Now we can use timers to implement timeouts. Timeouts that we were using earlier using Time.sleep. So instead of using Time.sleep this time we are going to use timer after. So this is another mechanism of using timers and it is also part of the time package. So let's explore time dot after the after method of time package. Now let's use Time.after for its timeout functionality. So first thing that I'm going to do is I'm going to create a function which will simulate a long running, which will simulate a long running operation, a long running operation, no arguments needed and no return values. So I'm going to create a loop and it's going to be ranging over 20. So 20 times. And I'm going to print the value of I. So it's going to be printing zero up to 19 after every loop iteration it's going to sleep. So time dot sleep and time dot second. So this function is going to. So this function is going to be running for 20s. And we have used time dot sleep here. But that is only to simulate that this function is taking time. Maybe it's fetching data from somewhere. Maybe it's just fetching data from database. Or maybe it's fetching a big file that the user saved on our server. So it could be any situation, but it is a task that is heavy and it will take a considerable time to finish. Now let's come to our main function func main. And here we are going to use our long running operation. But first we are going to define a timeout. So timeout is going to be time dot after two seconds. So we know that this operation should usually take not more than two seconds. But in this scenario it's going to take more than two seconds. And we are going to handle that situation now. So give it a time duration of two seconds. And then we create another channel. And this channel is going to receive a completion Value that the task has been finished. This is how we use channels to send signals from Goroutines that a task has been finished, and then something needs to be done in the main function or in the corresponding function. So let's create this channel, make channel bool and now we create a goroutine. And in this goroutine we are going to call long running operation. And once the long running operation is finished, then we are going to send a signal that the work is done. So we send true because done is a boolean channel. Now we use select case because now we have a done channel. And let's see what Time.after After returns. So we are creating a variable timeout. And it is going to store value that is received from Time.after. So what type of value is that? It is a channel. It is a channel of type time dot time. So we use select case because we are handling channels. And channels are handled using select case. The case statement only accepts statements. So select and then case. The first case will be time out and then Fmt.println operation timed out. Let's define the next case as well. Case when we receive something from the done channel. All right. So Fmt.println operation completed. That's it. Let's see what happens. And here we go. We only have two iterations of this long running operation, because it could only perform this operation for two seconds. After that, it timed out. And that's why we have a message operation times out. It should be timed out. All right. So let's again run this. This time we use three. So we should see 012. Now. And now timed out. Okay. So this is how we are timing out functions. And we are using a time function that does not block the rest of the operations. The rest of the instructions in our code. Perfect. So now let's create another example. So let me comment this out the whole thing. Okay. So. Time out. So this is time out. And then this was. Basic timer use. And this time we are scheduling delayed delayed operations. So how do we do that. Let's start our main function now func main. But in lower case func main. All right let's initiate a timer. So timer. New timer. Time out, new timer. And we are going to give it a duration of two seconds. Two into timer second. And now a go routine. Go func. And immediate execution. And then we receive the value from timer dot c here. So timer dot capital C. So once the timer has expired then we are going to print a statement. It's going to be delayed. Operation executed. Delayed. Operation executed. Now let's just print Fmt.println and we are waiting. Perfect. So in the first example here, this was blocking in nature. And I told you that we can use this in a go routine and that's what we are doing. So when we do that we are scheduling a delayed operation. Let's see this in action. What happens here. Let's see what happens now. And then we'll discuss. So run this and here we go. So all we see is waiting printed out onto the screen. So why did that happen. Well that happened because Time.new timer is non-blocking in nature. We don't have any blocking mechanism to wait for the go routine, and hence we are not waiting for the go routine to come back to the main thread. This goes out. This doesn't block execution. Workflow comes here. It prints this string and then ends the program. Perfect. So now let's use time dot sleep. And we have to use this because we want to block the program timer functions Time.after or Time.new timer. They are non-blocking. So we have to use Time.sleep for blocking. So we use three seconds now because yes our timer is two seconds. So three into time dot second. And let's also add another print statement. And I'm going to type end of the program so that you see in action the difference between time dot sleep and time dot new timer. Let's scroll this a little and then clear this run this waiting Delayed operation executed and end of the program. So let's go over this program and see what happens. At this point. This timer starts. All right non non-blocking timer starts. All right. Now this is counting down or counting up whatever. It's counting two seconds. All right. So counting two seconds starts here. At this point this is a go routine because it has a go keyword. So this function executes and it leaves our main thread even before it executes. All right. So this function leaves our main thread and goes out back in another thread okay. And this. So we are not going to discuss this go routine here. So this has left the main thread. We are still in the main thread. Now we move on to the next line Fmt.println waiting. So we print waiting onto the terminal, and then the execution comes to the next line. The runtime comes to the next line. It sees time.sleep three into time dot second. So now this is blocking. So this is a blocking timer. And it starts here. So we are already using this timer. It started and it was counting two seconds. This blocking timer starts and it starts counting three seconds okay. Now here we are blocking. Now this finishes after two seconds and it still has one second remaining. So we still have time to send the time in. After two seconds it sends the current time to the channel. So after two seconds it has time to send the current time to the channel. All right. So it sends a time the current time to the timer dot See? And then as soon as it sends the time to the channel, we have a receiver ready to receive that time. As soon as this receiver receives some value, then it moves on to the next line, which prints delayed operation executed. It comes back to the main thread and prints delayed operation executed. Now after this, maybe one second, or maybe a little less than one second, maybe 900 milliseconds or something like that. So maybe just exactly one second, or maybe a little less than one second is remaining. So we wait for that once this is printed out and then end of program gets printed out. So now I'm going to run this again. Notice the execution timing of the output. So here we go. Perfect. You saw that after reading the pause was longer but delayed. Operation executed. Once this was printed, the pause was a little shorter. Now let's come to some best practices for using timers. The first recommended best practice for using timers by the go developers. The go official team is to always stop timers to avoid resource leaks. We need to always remember to stop timers when they are no longer needed to avoid resource leaks, and we should use defer to ensure proper cleanup. Right. So we can use defer and then timer dot stop. Even if the timer has expired we still need to stop the timer. So why is it important? It's important because we need to manage the life cycle of a timer properly to ensure efficient resource usage. We don't just code to get one thing done, we want to print something onto the screen and this is all we are concerned about. Know when we are coding. We are not just coding to get a desired output, but to get the desired output in the most efficient way. And we always need to make sure that we are using resources carefully. We are always monitoring resource usage. That is the quality of a good coder, not just a good coder, an experienced coder, a coder who has expertise in his work. So if you want to reach that level, then you need to start ensuring that you always monitor resource utilization. You are careful about resource usage and you are cleaning up resources as you are using them. So why are we stopping a timer? Well, stopping a timer helps in freeing up resources and avoiding unexpected behavior. If you do not stop a timer and it is no longer needed, it will still consume resources. Although the timers channel will eventually be garbage collected, the timer itself will remain until it either fires or is stopped. All right, let's close this and let's come to our final example where we use multiple timers. So let's comment this out and let's create this small example func main. Timer one. So time dot new timer. And then this will be for one second time dot one second, and then time dot and then two into time dot second. Timer two. Now you select case to handle both the timer channels. All right. So case first. Timer. Timer one dot c capital C and then Fmt.println. Timer one expired. Then case timer two dot C. Fmt.println. Timer two expired. And that's it. So this is just a very, very simple, basic example of handling multiple timers. But it can get very complex in no time. Let's run this. All right, so only one of the case needs to be fulfilled. And then the select case will come out of the select statement and end the program. Perfect. If you want to use both of them then we use a for loop which will iterate over this select case twice. All right. So I assume that you know by now how you can do that. So you can do that on your own in your program and run that. Overall in this lecture on timers in Go, we covered essential concepts practical examples for effectively using timers to handle time based operations. This lecture provides a thorough understanding of how to use timers for timeouts, delays, and periodic tasks, ensuring efficient and responsive Go applications.

*/
