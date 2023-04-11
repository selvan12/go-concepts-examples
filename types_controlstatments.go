package main

/* Go's basic data types are,
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

/*
Zero values
Variables declared without an explicit initial value are given their zero value.

The zero value is:
0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.
*/

/*
Constants
Constants are declared like variables, but with the const keyword.

Constants can be character, string, boolean, or numeric values.

Constants cannot be declared using the := syntax.

const Pi = 3.14
*/

/* For
Go has only one looping construct, the for loop.
The basic for loop has three components separated by semicolons:

the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration

For is Go's "while"
At that point you can drop the semicolons: C's while is spelled for in Go.
	for sum < 1000 {
		sum += sum
	}
*/

/* If
If with a short statement
Like for, the if statement can start with a short statement to execute before the condition.
	if v := math.Pow(x, n); v < lim {
		return v
	}
*/

/* Switch
A switch statement is a shorter way to write a sequence of if - else statements. It runs the first case whose value is equal to the condition expression.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
Switch without a condition is the same as switch true.
This construct can be a clean way to write long if-then-else chains.
*/

/* Defer
A defer statement defers the execution of a function until the surrounding function returns.
The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Stacking defers
Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
for i := 0; i < 10; i++ {
	defer fmt.Println(i)
}
*/
