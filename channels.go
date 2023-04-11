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

	rangeAndCloseChannel()
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

/*
Range and Close Channel:
A sender can close a channel to indicate that no more values will be sent.
Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

v, ok := <-ch
ok is false if there are no more values to receive and the channel is closed.

The loop for i := range c receives values from the channel repeatedly until it is closed.

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.
*/
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func rangeAndCloseChannel() {
	fmt.Println("rangeAndCloseChannel with fibonacci example")
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
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
quit
rangeAndCloseChannel with fibonacci example
0
1
1
2
3
5
8
13
21
34

*/

/*
Understanding Inner workings of the Golang Channels
***************************************************

ch := make(chan int, 3)
The statement above creates a buffered channel capable of holding up to 3 values of type int

Under the hood, the functionmake allocates an hchan struct on the heap and returns a pointer to it.

Here are some fields of the hchan struct and their explanations.

type hchan struct {
	buf      unsafe.Pointer
	sendx    uint
	recvx    uint
	lock     mutex

	...   // other fields
}

buf is a pointer to an array, which maintains a circular queue
sendx is the index of the sent element in the array
recvx is the index of the received element in the array
lock ensures that the reading and writing of the channel is an atomic operation


When the channel is not full, we can inserts elements at the back of the circular queue without block.
// G1 sends three elements into the channel, capicity = 3
ch <- elem1
ch <- elem2
ch <- elem3

When the channel is not empty, we can receive elements from the front of the circular queue without a block.
// G2 receive three elements from the channel, capicity = 3
<- ch
<- ch
<- ch


Waiting Goroutines
Other fields are also important when handling blocking between goroutines.

recvq stores the blocked goroutines while trying to read data on the channel.
sendq stores the blocked goroutines while trying to send data from the channel.
Remember that Both of recq and sendq are linked list.

type hchan struct {
 buf      unsafe.Pointer
 sendx    uint
 recvx    uint
 lock     mutex

 sendq    waitq
 recvq    waitq
 ...   // more fields
}
type waitq struct {
  first *sudog
  last  *sudog
}
// pseudo goroutine
type sudog struct {
  g     *g
  elem  unsafe.Pointer
  next  *sudog
  prev  *sudog
  ...
  c     *hchan
}

When the channel is empty, a receive operation leads to the blocking of the current goroutine. All the blocked goroutines are stored inside the recvq queue.

So, when does the blocked goroutine get resumed?
The answer is when a new goroutine performs a send operation on the channel.

Here are the details.
A new goroutine copied the new data directly into the first waiting goroutine’s element
The first waiting goroutine is popped off from recvq
Runtime scheduler set the popped goroutine runnable and put it on its’ runqueue. Then the blocked goroutine is triggered and ready to run again.

When the channel is full, the next send operations block their respective goroutines. All the blocked goroutines are stored inside sendq queue.

Until another goroutine’s receive, the blocked goroutine is resumed. Here are the details:
When a new goroutine performs a receive operation on the channel, the first element in the buffer is removed
The first waiting goroutine is pop off from sendq
The element of the poped goroutine is copied into the buffer
Runtime scheduler set the popped goroutine runnable and put it on its’ runqueue. Then the blocked goroutine is triggered and ready to run again.


Reference:
https://levelup.gitconnected.com/how-does-golang-channel-works-6d66acd54753
*/
