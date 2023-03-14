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

The evaluation of f, parameters(s here) happens in the current goroutine and the execution of f happens in the new goroutine.

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
