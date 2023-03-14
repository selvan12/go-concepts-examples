package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
/* Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
(The data flows in the direction of the arrow.)
Internally, works like a FIFO circular queue

Create a new channel with make(chan val-type). Channels are typed by the values they convey.
Send a value into a channel using the channel <- syntax.
The <-channel syntax receives a value from the channel.
Like maps and slices, channels must be created before use:

ch := make(chan int)
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
*/
func channelExample() {
	fmt.Println("Start")

	msg := make(chan string)

	go func() { msg <- "test chan" }()
	msgOut := <-msg
	fmt.Println(msgOut)

	fmt.Println("Another Example")
	slice := []int{7, 9, 4, -11, 1, 0}
	count := make(chan int)
	fmt.Println("First half values ", slice[:len(slice)/2])
	fmt.Println("Second half values ", slice[len(slice)/2:])
	go sumMembers(slice[:len(slice)/2], count)
	go sumMembers(slice[len(slice)/2:], count)
	x, y := <-count, <-count
	fmt.Println("Values are: ", x, y, x+y)

	bufferedChannel()

	selectExample()
}

func sumMembers(slice []int, count chan int) {
	sum := 0
	for _, val := range slice {
		sum += val
	}
	count <- sum
}

// Buffered Channels
/* Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

By default channels are unbuffered, meaning that they will only accept sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value.
Buffered channels accept a limited number of values without a corresponding receiver for those values.
*/
func bufferedChannel() {
	fmt.Println("Buffered Channel Example")
	buffChan := make(chan string, 2)
	buffChan <- "buffer 1"
	buffChan <- "buffer 2"
	fmt.Println(<-buffChan)
	fmt.Println(<-buffChan)
}

// The select statement lets a goroutine wait on multiple communication operations.
/* A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

Go’s select lets you wait on multiple channel operations. Combining goroutines and channels with select is a powerful feature of Go.

The default case in a select is run if no other case is ready.
Use a default case to try a send or receive without blocking.
*/
func selectExample() {
	fmt.Println("Select Example")
	chan1 := make(chan string)
	chan2 := make(chan string)
	quit := make(chan string)
	go func() {
		time.Sleep(time.Second * 1)
		chan1 <- "chan1"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		chan2 <- "chan2"
		quit <- "quit"
	}()
	//quit <- "quit"

	for {
		select {
		case msg := <-chan1:
			fmt.Println(msg)
		case msg := <-chan2:
			fmt.Println(msg)
		case msg := <-quit:
			fmt.Println(msg)
			return
		default:
			time.Sleep(time.Millisecond * 500)
			fmt.Println(".")
		}
	}
}

/* Output:
% go run main.go channels.go
2023/03/12 17:11:57 Welcome to Go Concepts examples
Start
test chan
Another Example
First half values  [7 9 4]
Second half values  [-11 1 0]
Values are:  -10 20 10
Buffered Channel Example
buffer 1
buffer 2
Select Example
.
.
.
chan1
.
chan2
.
quit
*/
