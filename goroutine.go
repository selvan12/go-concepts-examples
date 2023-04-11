package main

import (
	"fmt"
	"log"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime.
/* Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.
To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
You can also start a goroutine for an anonymous function call.

The evaluation of f and parameters(s here) happens in the current goroutine and the execution of f happens in the new goroutine.

Goroutines run in the same address space, so access to shared memory must be synchronized.
The sync package provides useful primitives, although you won't need them much in Go as there are other primitives.
*/
func routineExample(msg string) {
	for i := 0; i < 3; i++ {
		fmt.Println(msg, " : ", i)
	}
}

func goRoutineExample() {
	log.Println("Start")

	routineExample("func")

	go routineExample("routine")

	go func(msg string) {
		fmt.Println(msg)
	}("another routine")

	time.Sleep(time.Second * 2)
	log.Println("Done")
}

/* Output:
go run main.go goroutine.go
2023/03/12 13:57:13 Welcome to Go Concepts examples
2023/03/12 13:57:13 Start
func  :  0
func  :  1
func  :  2
another routine
routine  :  0
routine  :  1
routine  :  2
2023/03/12 13:57:15 Done
*/

/*
How Goroutine actually work:
***************************
One of the main reasons that the Go Language has gained incredible popularity in the past few years is
the simplicity it deals with concurrency with its lightweight goroutines and channels.

Goroutines are the way of doing tasks concurrently in golang.
They exist only in the virtual space of the Go runtime and not the OS, therefore the Go Runtime scheduler is needed to manage their lifecycles.

Go Runtime maintains three C structs for this purpose: (https://golang.org/src/runtime/runtime2.go)

The G Struct : Represents a single goroutine and contains the fields necessary to keep track of its stack and current status. It also contains references to the code that it is responsible.
The M Struct : Represents an OS thread. It also contains pointers to fields such as the global queue of runnable goroutines, the current running goroutine, its own cache and the reference to the scheduler
The Sched Struct : It is a single, global struct that keeps track of the different queues of goroutines and M's and some other information that the scheduler needs in order to run, such as the Global Sched Lock.

There are 2 queues containing G structs, 1 in the runnable queue where M's (threads) can find more work, and the other is a free list of goroutines.
There is only one queue pertaining to M's (threads) that the scheduler maintains. And in order to modify these queues, the Global Sched Lock must be held.

So, on startup, go runtime starts a number of goroutines for GC (Garbage Collector), scheduler and user code. An OS Thread is created to handle these goroutines.
These threads can be at most equal to GOMAXPROCS (This is defaulted to 1, but for best performance is usually set to the number of processors on your machine).

To make the stacks small, Go’s run-time uses resizable, bounded stacks, initially of only 2KB/goroutine.
A newly minted goroutine is given a few kilobytes, which is almost always enough. When it isn’t, the run-time grows (and shrinks) the memory for storing the stack automatically, allowing many goroutines to live in a modest amount of memory.

The Go Runtime Scheduler keeps track of each goroutine, and will schedule them to run in turn on a pool of threads belonging to a process.

The Go Runtime Scheduler does cooperative scheduling, which means another goroutine will only be scheduled if the current one is blocking or done, and that is easily done via code. Here are some examples:

Blocking syscalls like file and network operations.
After being stopped for garbage collection cycle.


Reference:
https://medium.com/the-polyglot-programmer/what-are-goroutines-and-how-do-they-actually-work-f2a734f6f991
*/
