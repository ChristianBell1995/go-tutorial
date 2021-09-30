package main

import (
	"fmt"
)

func main() {
	// Channels are the pipes that connect concurrent goroutines. You can send values into channels
	// from one goroutine and receive those values into another goroutine.
	// In unbuffered channels the sender and receiver of the channel must be "ready" -
	// in other words there is no capacity to store values that are passed into the channel.

	// Example 1
	intChan := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			intChan <- i
			fmt.Println("Successfully sent %v", i)
		}
		// close(intChan)
	}()

	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan)

	// for elem := range intChan {
	// 	fmt.Println(elem)
	// }

	// Make an unbuffered channel deadlock!
	myChan := make(chan string)
	// go func() {
	// 	fmt.Println(<-myChan)
	// }()
	myChan <- "hi"

	// For unbuffered channels there must be a sender and a receiver ready otherwise you get a deadlock!

	// Buffered channels accept a limited number of values without a corresponding receiver for those values. So you can push values in whilst outside a goroutine.
	// bufferedIntChan := make(chan int, 2)
	// bufferedIntChan <- 1
	// bufferedIntChan <- 2
	// bufferedIntChan <- 3
	// fmt.Println(<-bufferedIntChan)
	// fmt.Println(<-bufferedIntChan)

	// Using channels to block execution of threads:
	// boolChan := make(chan bool)
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	boolChan <- true
	// }()
	// fmt.Println("Waiting for timer")
	// <-boolChan
	// fmt.Println("Timer has finished!")
}
