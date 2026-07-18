package advanced
// 012 Tickers

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

/*
A tiger in golf is a mechanism for producing ticks at regular intervals. And when I talk about producing ticks, I'm not talking about little insects. Tickers are useful for performing periodic tasks or operations on a consistent schedule. Tickers are often used in scenarios where tasks need to be repeated at fixed intervals, such as polling, periodic logging, or regular updates. Using tickers ensures operations occur at regular intervals, maintaining a consistent schedule. Using tickers also simplifies the implementation of recurring tasks without manually handling timing logic. Ticker is created using Time.new ticker, so we are using time package and with that we use new ticker function to create a new ticker with a specific time interval, similar to timers, tickers also have a channel associated. So when we create an instance of a ticker, it will have a key field of the ticker type, which is a channel that receives ticks at regular intervals. Now let's see our ticker in action. So let's create a ticker using time dot new ticker. And the only argument is time duration. We are going to give it time dot second. That's it. And now we use a for loop to range over the ticker. So for now I'm going to use the values that this ticker returns for demonstration purposes. So Fmt.println. And let me include a helping string. So tick at and we use tick value. All right. Okay I forgot to use range keyword so far. Range. Sorry about that. Now let's save this and go run ticker and go run tickers. So tick tick tick and it'll keep on ticking. All right. So this is an infinite loop. And we have given it a duration of time dot second. So that's why this for loop is looping over this range of values after every second. Because it gets populated after every second. Now all we need to do is hit control. Control C and that's it. Now, if you were using Tiger dot C for something else and not using it to print time after every second, we would use it something like this. So for range ticker dot C and I'm just going to create a simple loop where we increase value of a variable by one. So I'm going to introduce I as zero here. And then I plus plus and let's print the value of I. Let's run this now. So one two. So every second the value of I is increased. And we get to see that on the terminal. So instead of I it could be anything that you want to implement it. And instead of time dot second it could be two multiplied by time dot second. And we keep multiplying it by two. So We do this and save it. Stop this and run it again. So now after every two seconds, I will have. Okay, I won't because I is zero. Make it one and then save it. Run it again. And that's how it will keep moving forward. So let's stop this. And now we can see that we are interrupting our program to stop the ticker. How do we stop this ticker in our program? Similar to timers, we have a stop method to stop a ticker. It is important to stop a ticker to release resources and prevent it from producing further ticks. Because ticker does not expire. Timers have an expiry, but tickers Don't. And that's why we always have to stop our ticker use differ with stop to stop the ticker before the end of the functions return. So let's do that here. Differ ticker dot stop and and instead of ranging over ticker dot C because that is infinite loop that's range over five. So after five times we will stop the ticker because as soon as the main function returns, it will not return without executing the differ statements. So before returning before ending the main function, differ will execute and it will stop the ticker and hence we will not leave the ticker in a limbo in our computer's memory. Right? So let's run this now. clear this and run this. Here we go. So this time we did not have to use Ctrl C to stop the program. So what's the practical use for tickers? Well, a very common use for tickers is periodic task execution. We can use a ticker to execute tasks at regular intervals, such as polling data, updating status, or performing routine maintenance. Let's see how we do that. So let's close the terminal. And I'm going to comment this out. And actually one moment instead of writing over the range ticker dot c I'm going to keep it here. Tick and. Tick. And then comment it. and then and then completely comment this and this as well. All right. Perfect. Now let's create a different function where we will practice a periodic task execution. So let's first create the periodic task. So periodic task. And I'm going to just print something in this periodic task. So performing periodic task at this is a string. And then just time. So time dot. Now I'm going to print the current time. Now let's come to our main function. And we first introduce a ticker Time dot new ticker. Give it a duration of 10s. Actually, 10s. We'll have to wait for too long. So. Time dot second and one second. We don't need to mention one then. Now let's close this. So ticker dot stop. And again an infinite loop with select case. So we can use range as well. But we haven't used select case with ticker. So just for variety I'm going to use select case. So select and then case. Ticker dot see. And then we and then we run periodic task. So let's run this now. And after every second we are performing a periodic task. Now, this is just for example, sake. It could be an hour. So instead of time. Second, we could be using an hour as a duration. And after every hour, let's say we are cleaning up some operations, or we are sending some data over to our database or over to our event management pipeline, or to our logging mechanism, or we are performing some bulk logging operations and it could be anything. There are n number of cases that we can perform at periodic time intervals in our applications, in our APIs. So this is what it is for. And again, if there is an error, if for some reason our main function panics, or if there is an error that stops our application or our API, then this ticker should not linger on in the memory. So if something like that happens, then differ will be executed and the ticker will be stopped. And this kept on going because it is a real time scenario where we want our application to continue running. If it's an API, we want our API to continue running and hence the periodic task will always be running. There will be no expiry because our API has no expiry. So this was an example of periodic task execution. Let's now move to polling for updates. So this is a very basic example of running a periodic task. Or if we are polling for updates, we are if we are implementing a polling mechanism that periodically checks for updates or changes. So that's also going to be implemented in the same way. And if we are scheduling any logging mechanism after a certain period of time, then that's also going to be implemented in exactly the same way. We will define a function with logging mechanism, and then we will start a ticker and give it a duration. And after that duration our function will keep on executing. So it will be the same mechanism. A basic structure of that mechanism will be same as this. So whether we are scheduling a logging mechanism or polling for resources, or we are running a periodic task, it's going to be working in the same way. Now what do we do if we want to stop the ticker and print a message or perform an action once the ticker stops? How do we handle that? How do we create that mechanism so we don't just have to rely on ticker by now. We also know how to use timers. So we are going to use an amalgamation of ticker and timer. So. First let me create a heading here. So scheduling logging periodic tasks or polling for polling for updates. So this can be for scheduling logging periodic tasks or polling for updates. Perfect. Now let's handle our ticker. Stopping gracefully. So when we handle ticker stops gracefully, that means that we can either print an ending statement or we can execute a function, or we can do anything when the ticker stops. So func main. And first we create a ticker time dot new ticker and give it time dot second that's it. And now we have a stopper. So stop. And this is going to be timed out after because time after returns a channel. All right. So we remember that timed out after returns a channel. So we can implement a select case where we can handle the ticker channel. And then once the timer stops this time after once it expires it will return us a channel. And we can handle the return value within that select case statement. So let's give it five seconds. So five multiplied by time dot second. And now we always need to defer ticker dot stop. Perfect. Now let's introduce an infinite loop because we don't have to worry about it. Timer is going to stop the timer after, and once it stops, we are going to return from the function. So first we start with select and the first condition is that we receive value from ticker dot c. So we receive values using the receive operator. The arrow operator. All right. Now let's use tick to print a statement. So fmt.println and we'll mention tick at and then the tick value. Next we have another case. And that is the stop channel. Stop is actually a channel Time.after returns us a channel of time dot time. So this is a channel. So we just use a receiver on this channel. A receiver operator is arrow before the channel. So arrow. And then after that we mention the channel which is top. And then we handle return. But before that we are going to print a statement. So fmt.println stopping ticker. And instead of just printing something out you can customize your implementation. Maybe you need to run further functions after the ticker stops. Maybe you need to use these values. Maybe you need to run some more functions. Once the ticker stops using these ticker values or using some other values. So this is just basic, but we can customize it to our own will. So save this and clear this and then go run tickers. Perfect. So we are getting the output. And now it stopped and we are stopping the ticker gracefully. We are sending a message to the user that we are stopping ticker now. So this was a good example where we used ticker along with timer. So this is a ticker pattern that you may encounter in many real world scenarios. Because there are many scenarios where we combine tickers with timers to create complex timing logic, where certain tasks need to be performed periodically with timeouts or delays. Similarly, you can handle multiple tickers as well. However, in this lecture I'm not going to take up multiple tickers. What I'm going to do is I've never given you a task, so I'm going to ask you to create your own example of multiple tickers and paste it in the comment section in this lecture. So practice multiple tickers on your own. But always remember the basic thing when we are creating a ticker, we have to stop the ticker. All right. So don't forget to stop the ticker. Defer the stop and then continue with your code. By now you must have noticed that I don't give any assignments. And I try to cover as many things as I can in the classroom in our lectures, so that you can understand as many things as possible. You can get to practice as many things as possible with me. It may be possible that some people may leave some topics if they are not discussed in classroom in our lectures, and I don't want that to happen. And that's why I try to be as detailed in my explanation and as thorough with implementation of any package or any feature of the language. However, I would still ask you to keep experimenting with the topics that we have discussed in any lecture, every lecture that you have finished, make it a point that in your free time, some day, maybe later in the day, you experiment on your own in that topic. Experiment different scenarios because that is the most fun part. This part of learning different topics in any language is the most fun part. Once we start making APIs, once you start making applications and APIs in the language, it's not as fun. But this this is fun because you get to experiment and notice different things, different outcomes. You make mistakes when you create your own scenarios, the code may run into a problem, and then you will realize the behavior of that feature, behavior of that statement, behavior of that loop, or whatever the topic was of that lecture. You will notice the behavior of that specific thing in the language in a certain scenario that you executed, and then you will correct that scenario. You will correct something and then it will function perfectly. So you will notice something new that may not have been covered in the lecture, because any book or any course that you take, programming languages are so, so deep that you cannot cover everything in any book or any course. So when you start experimenting with the features of the language, then you will start exploring newer things. And trust me, you will have fun. It will be very, very fun because once you start making applications and APIs then yeah, it's it's it's a different ball game altogether. You you have a task and you need to get that task. It's still I mean making an API and applications are also a learning based activity. But here we have all the aspects of the programming language that we are experimenting with, and we are making short programs, we are making smaller programs and we are observing the behavior in different scenarios. So that's why it is different from when we are working on an API, because that's a that's a bigger code base, and it's a complex code structure where there are limitations on how much we can experiment. So enjoy this time and utilize this time. All right. So even if you don't have any assignments create your own assignments. Right. You need to maintain your own discipline of practicing things on your own instead of somebody else giving you an assignment. Give yourself some assignments. Give yourself some tasks. Assign some time for some personal practice. In conclusion, tickers in go are powerful tools for handling periodic tasks and operations, providing an efficient way to execute code at regular intervals. This comprehensive lecture has explored the fundamental concepts of tickers, provided practical examples, and highlighted best practices and advanced patterns to help you use tickers effectively in your go. Applications. By mastering tickers, you can improve the reliability and performance of your time based operations, ensuring that your applications run smoothly and efficiently. Understanding and utilizing tickers is crucial for developers who need to manage periodic tasks, implement recurring events, or maintain consistent intervals in their go programs. With the knowledge gained from this lecture, you are now equipped to integrate tickers into your applications, confidently enhancing your code's functionality and robustness.
*/
