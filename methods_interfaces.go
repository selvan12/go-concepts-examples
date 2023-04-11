package main

/* Methods
Go does not have classes. However, you can define methods on types.
A method is a function with a special receiver argument.
The receiver appears in its own argument list between the func keyword and the method name.
In this example, the Abs method has a receiver of type Vertex named v.
	func (v Vertex) Abs() float64 {
	}
Remember: a method is just a function with a receiver argument.

You can declare a method on non-struct types, too.
In this example we see a numeric type MyFloat with an Abs method.
	type MyFloat float64
	func (f MyFloat) Abs() float64
You can only declare a method with a receiver whose type is defined in the same package as the method.
You cannot declare a method with a receiver whose type is defined in another package (which includes the built-in types such as int).

You can declare methods with pointer receivers.
This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
Methods with pointer receivers can modify the value to which the receiver points. Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
*/

/* Interfaces
An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.

A type implements an interface by implementing its methods.
	type I interface {
		M()
	}

	type T struct {
		S string
	}
This method means type T implements the interface I,
but we don't need to explicitly declare that it does so.
	func (t T) M() {
		fmt.Println(t.S)
	}

The empty interface
The interface type that specifies zero methods is known as the empty interface:
	interface{}
An empty interface may hold values of any type.
*/

/* Type assertions
A type assertion provides access to an interface value's underlying concrete value.
	t := i.(T)
This statement asserts that the interface value i holds the concrete type T and assigns the underlying T value to the variable t.
If i does not hold a T, the statement will trigger a panic.
To test whether an interface value holds a specific type, a type assertion can return two values: the underlying value and a boolean value that reports whether the assertion succeeded.
	t, ok := i.(T)
If i holds a T, then t will be the underlying value and ok will be true.
If not, ok will be false and t will be the zero value of type T, and no panic occurs.
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)
*/

/* Type switches
A type switch is a construct that permits several type assertions in series.
A type switch is like a regular switch statement, but the cases in a type switch specify types (not values), and those values are compared against the type of the value held by the given interface value.

	switch v := i.(type) {
	case T:
		// here v has type T
	case S:
		// here v has type S
	default:
		// no match; here v has the same type as i
	}
The declaration in a type switch has the same syntax as a type assertion i.(T), but the specific type T is replaced with the keyword type.

func do(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("Twice %v is %v\n", v, v*2)
		case string:
			fmt.Printf("%q is %v bytes long\n", v, len(v))
		default:
			fmt.Printf("I don't know about type %T!\n", v)
	}
}
func main() {
	do(21)
	do("hello")
	do(true)
}
*/

/* Stringers
One of the most ubiquitous interfaces is Stringer defined by the fmt package.

	type Stringer interface {
		String() string
	}
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.
	type Person struct {
		Name string
		Age  int
	}

	func (p Person) String() string {
		return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
	}

	func main() {
		a := Person{"Arthur Dent", 42}
		z := Person{"Zaphod Beeblebrox", 9001}
		fmt.Println(a, z)
	}
*/

/* Errors
Go programs express error state with error values.
The error type is a built-in interface similar to fmt.Stringer:

	type error interface {
		Error() string
	}
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)
Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

	i, err := strconv.Atoi("42")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
		return
	}
	fmt.Println("Converted integer:", i)
A nil error denotes success; a non-nil error denotes failure.
*/
