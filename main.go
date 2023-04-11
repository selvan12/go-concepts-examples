package main

import "log"

/*
Advantages of Using Go
=============================

1. Simplified code
Golang code is less complex and borrows several concepts from other programming languages.
This comes with plenty of benefits – developers need to traverse fewer lines of code, saving a ton of time.
This is important since writing code takes less time than actually reading it.

2. Powerful performance
The simplicity of Go adds up to its powerful performance.
It runs faster, compiles quicker, it’s easy to maintain and support, and allows for shorter software development lifecycles.
Golang compiles directly to machine code as it doesn’t use Virtual Machines, providing an even better speed advantage.

3. The go-to language for large-scale projects
Many enterprises choose Golang because it accelerates and increases work effectiveness on more challenging projects.
Golang is based on the concept that, preferably, there should be only one solution to any given problem.
In contrast to other languages, where there may be virtually as many solutions as programmers, making project collaboration even more challenging.

4. Designed for multi-core processors
Go is designed for the modern era of cloud computing and parallel processes utilized by the current processors.
Other popular computer programming languages like Java, JavaScript, Python, Ruby, or C, C ++ were created before the multi-core computers were widely used.
With Golang, it is easier to utilize all CPU cores without complicating development.

5. Designed for the internet
With the Golang standard library, developers can build complex web services without third-party libraries.
It makes it perfect for web development with fewer iterations and fewer possibilities to implement the same feature.

6. Fast garbage collection
Any application’s performance is greatly enhanced when you have a quick garbage collector like in Go.
The system’s memory isn’t crammed with junk supporting the development of fast-running apps.

7. Easy to learn
Programmers frequently see Golang as C without the inconvenient flaws.
Its syntax appears familiar to developers, making it easier to learn if you are familiar with the C concepts.

8. Easy maintenance
One benefit noted by businesses utilizing Go is the abundance of tools available to developers for automatic code maintenance.
The results are virtually identical to those produced by human programmers.

9. Open Source approach
Due to its open-source nature, Golang has many fans among programmers and a massive community around it.
Through forums, tutorials, and open-source projects, Go encourages the use of novel solutions while the community swiftly finds and fixes coding flaws.

10. Golang is industry agnostic
Adoption of the Go programming language may result in quicker and more productive coding, which boosts output.
This is why Golang is recognized by many big names across various industries.
=============================
*/

func main() {
	log.Println("Welcome to Go Concepts examples")

	//goRoutineExample()

	channelExample()

	//httpClient()
}

/*
An overview of memory management in Go
**************************************
As programs run they write objects to memory. At some point these objects should be removed when they’re not needed anymore. This process is called memory management.

Go offer automatic dynamic memory management, or more simply, garbage collection. Languages with garbage collection offer benefits like:
- increased security
- better portability across operating systems
- less code to write
- runtime verification of code
- bounds checking of arrays

A running program will store objects in two memory locations, the heap and the stack. Garbage collection operates on the heap, not the stack.
The stack is a LIFO data structure that stores function values. Calling another function from within a function pushes a new frame onto the stack, which will contain the values of that function and so on. When the called function returns, its stack frame is popped off the stack.

In contrast, the heap contains values that are referenced outside of a function.
For example, statically defined constants at the start of a program, or more complex objects, like Go structs. When the programmer defines an object that gets placed on the heap, the needed amount of memory is allocated and a pointer to it is returned.

Garbage Collection in Go
************************
Go prefers to allocate memory on the stack, so most memory allocations will end up there. This means that Go has a stack per goroutine and when possible Go will allocate variables to this stack.
The Go compiler attempts to prove that a variable is not needed outside of the function by performing escape analysis to see if an object “escapes” the function.
If the compiler can determine a variables lifetime, it will be allocated to a stack. However, if the variable’s lifetime is unclear it will be allocated on the heap.
Generally if a Go program has a pointer to an object then that object is stored on the heap.

Garbage collectors have two key parts, a mutator and a collector.
The collector executes garbage collection logic and finds objects that should have their memory freed.
The mutator executes application code and allocates new objects to the heap. It also updates existing objects on the heap as the program runs, which includes making some objects unreachable when they’re no longer needed.

The implementation of Go’s garbage collector
********************************************
Go’s garbage collector is a non-generational concurrent, tri-color mark and sweep garbage collector. Let’s break these terms down.

The generational hypothesis assumes that short lived objects, like temporary variables, are reclaimed most often. Thus, a generational garbage collector focuses on recently allocated objects.
However, as mentioned before, compiler optimisations allow the Go compiler to allocate objects with a known lifetime to the stack. This means fewer objects will be on the heap, so fewer objects will be garbage collected.
This means that a generational garbage collector is not necessary in Go. So, Go uses a non-generational garbage collector. Concurrent means that the collector runs at the same time as mutator threads. Therefore, Go uses a non-generational concurrent garbage collector.
Mark and sweep is the type of garbage collector and tri-color is the algorithm used to implement this.

A mark and sweep garbage collector has two phases, unsurprisingly named mark and sweep. In the mark phase the collector traverses the heap and marks objects that are no longer needed.
The follow-up sweep phase removes these objects. Mark and sweep is an indirect algorithm, as it marks live objects, and removes everything else.

Reference:
https://medium.com/safetycultureengineering/an-overview-of-memory-management-in-go-9a72ec7c76a8
*/
