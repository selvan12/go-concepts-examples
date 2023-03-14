package main

/* Pointers
Go has pointers. A pointer holds the memory address of a value.
The type *T is a pointer to a T value. Its zero value is nil.
	var p *int
The & operator generates a pointer to its operand.
	i := 42
	p = &i
The * operator denotes the pointer's underlying value.
	fmt.Println(*p) // read i through the pointer p
	*p = 21         // set i through the pointer p
This is known as "dereferencing" or "indirecting".
Unlike C, Go has no pointer arithmetic.
*/

/* Structs
A struct is a collection of fields.
Struct fields are accessed using a dot.
Struct fields can be accessed through a struct pointer.
A struct literal denotes a newly allocated struct value by listing the values of its fields.
*/

/* Arrays
The type [n]T is an array of n values of type T.
The expression,
	var a [10]int
declares a variable a as an array of ten integers.
An array's length is part of its type, so arrays cannot be resized.
*/

/* Slices
An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.
The type []T is a slice with elements of type T.
A slice is formed by specifying two indices, a low and high bound, separated by a colon:
	a[low : high]
This selects a half-open range which includes the first element, but excludes the last one.
The following expression creates a slice which includes elements 1 through 3 of a:
	a[1:4]

A slice does not store any data, it just describes a section of an underlying array.
Changing the elements of a slice modifies the corresponding elements of its underlying array.
Other slices that share the same underlying array will see those changes.
*/

/* Slice literals
A slice literal is like an array literal without the length.
This is an array literal:
	[3]bool{true, true, false}
And this creates the same array as above, then builds a slice that references it:
	[]bool{true, true, false}
*/

/* Slice defaults
When slicing, you may omit the high or low bounds to use their defaults instead.
The default is zero for the low bound and the length of the slice for the high bound.
For the array
	var a [10]int
these slice expressions are equivalent:
	a[0:10]
	a[:10]
	a[0:]
	a[:]
*/

/* Slice length and capacity
A slice has both a length and a capacity.
The length of a slice is the number of elements it contains.
The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.
The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

The zero value of a slice is nil.
A nil slice has a length and capacity of 0 and has no underlying array.
*/

/* Creating a slice with make
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
The make function allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:
	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
*/

/* Appending to a slice
It is common to append new elements to a slice, and so Go provides a built-in append function.
	func append(s []T, vs ...T) []T
The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.
The resulting value of append is a slice containing all the elements of the original slice plus the provided values.
If the backing array of s is too small to fit all the given values a bigger array will be allocated. The returned slice will point to the newly allocated array.
*/

/* Range
The range form of the for loop iterates over a slice or map.
When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
	for i, v := range pow {
	}
You can skip the index or value by assigning to _.
	for i, _ := range pow
	for _, value := range pow
If you only want the index, you can omit the second variable.
	for i := range pow
*/

/* Maps
A map maps keys to values.
	m := make(map[string]int)
The zero value of a map is nil. A nil map has no keys, nor can keys be added.
The make function returns a map of the given type, initialized and ready for use.

Map literals
Map literals are like struct literals, but the keys are required.
If the top-level type is just a type name, you can omit it from the elements of the literal.
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}
*/

/* Mutating Maps
Insert or update an element in map m:
	m[key] = elem
Retrieve an element:
	elem = m[key]
Delete an element:
	delete(m, key)
Test that a key is present with a two-value assignment:
	lem, ok = m[key]
If key is in m, ok is true. If not, ok is false.
If key is not in the map, then elem is the zero value for the map's element type.
*/
