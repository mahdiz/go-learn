# Learning Go

## General
Composite types are array, struct, pointer, function, interface, slice, map, and channel types—may be constructed using type literals.

```go
func swap(x, y string) (string, string) {
	return y, x
}
```

```go
func main() {
	var i, j = 1, 2
	var k, str, isFinished = 1, "foo", false
	a, b := swap("hello", "world")
	fmt.Println(i, j, k, str, a, b)
}
```

A defer statement defers the execution of a function until the surrounding function returns:

```go
func main() {
	defer fmt.Println("world")
	defer fmt.Println("my")
	fmt.Println("hello")
}
```

Deferred function calls are pushed onto a stack. This will output: “Hello my world”

## Loops
Go has only one looping construct, the for loop:

```go
for i := 0; i < 10; i++ {
	sum += i
}
```

While is similar to for loop:

```go
for sum < 1000 {
	sum += sum
}
```

Infinite loop:

```go
for {
}
```

Local variable in if block:

```go
if v := math.Pow(x, n); v < lim {
	return v
}
return v	// ERROR: v is undefined!
```

## Switch

```go
switch os {
case "darwin":
	fmt.Println("OS X.")
case "linux":
	fmt.Println("Linux.")
default:
	// freebsd, openbsd,
	// plan9, windows...
	fmt.Printf("%s.", os)
}
```

Switch with no condition is good to write long if-then-else chains:

```go
switch {
case t.Hour() < 12:
	fmt.Println("Good morning!")
case t.Hour() < 17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")
}
```

## Basic Types
```go
bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr byte (alias for uint8), rune (alias for int32), float32, float64, complex64, complex128
```

```go
var i int = 42
 var f float64 = float64(i)	// cast: T(v) converts the value v to the type T
```
NOTE: Unlike C, in Go assignment between items of different type requires an explicit conversion.

```go
const Pi = 3.14
```

## Pointers
Go has pointers similar to the behaviors of * and & in C. Unlike C, Go has no pointer arithmetic. 

```go
var i = 56
 var p *int = &i
 fmt.Println(*p)
```

## Arrays

```go
var a [3]string		// to define an array
primes := []int{2, 3, 5}	// to define and init at the same time
a[0] = "Hello"			// to access and set
var i = len(a)			// to get array length
```

## Concurrency 
A "go" statement starts the execution of a function call as an independent concurrent thread of control, or goroutine, within the same address space:

```go
// nameless goroutine
go func(x, y) {
	panic("Signal: ")
}()

// named goroutine
func findSum(x int, y int) {
	...
} 

go findSum(x, y)
```

### Channels
Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

```go
channel := make(chan bool)

go func() {
	time.Sleep(2 * time.Second)
	channel <- 'ping'
}()

fmt.Println("Goroutine started...")

// Wait until something is written on the channel
x := <-channel

fmt.Println("Goroutine finished!")
```

## Object-Oriented
Go does not support all object-oriented mechanisms common in Java and C#, and it is for a reason: keeping objects lightweight. Go does not support type inheritance, and its subclassing and polymorphism is limited.

A `struct` is a collection of fields: 

```go
type Vertex struct {
	x int
	y int
}

func (v *Vertex) addVertex(v2 Vertex) {
	v.x += v2.x;
	v.y += v2.y;
}

var v Vertex		// to create a Vertex
v := Vertex{56, 31}	// to create and init at the same time
v := Vertex{x: 56}	// to create and set specific fields
fmt.Println(v.x)		// to access fields of a struct
```

The (v *Vertex) is called a _method receiver_. If the method receiver is not defined as a pointer, then the object will be copied by value, i.e., changes to the object will not remain after the method returns. For example, the following code will print `{1,2}` while it would print `{4,6}` if we had `func (v *Vertex)`:

```go
func (v Vertex) addVertex(v2 Vertex) {
	v.x += v2.x;
	v.y += v2.y;
}

v1 := Vertex{x: 1, y: 2}
v2 := Vertex{x: 3, y: 4}
v1.addVertex(v2)
fmt.Println(v1)
```

Inheritance is done using embedding which is to embed a `struct` name inside another `struct`:

```go
type animal struct {
}

type duck struct {
	animal
	featherCount int
}
```

Go does not support [type inheritance](https://golang.org/doc/faq#inheritance). So, something like this won't compile:
 
```go
var a animal
d := duck{}
a = d		// COMPILE ERROR: cannot use d (type duck) as type animal in assignment 
```

Instead, we can use interfaces. They are satisfied implicitly, unlike Java or C#, so interfaces can be defined for code we don’t own.

```go
type bird interface {
	Fly() string
}

func (d *duck) Fly() {
	return "Duck flying..."
}
```

Now, we can write:

```go
var b bird
d := duck{}
b = d
```

### Polymorphism
Polymorphism in Go is limited and can only be achieved through interfaces. Method `Run()` in the following code has a polymorphic behavior:

```go
type ClientProtocol struct {
	Id int
}

type ServerProtocol struct {
	IsUp bool
}

type Protocol interface {
	Run() error
}

func (p ClientProtocol) Run() error {
	fmt.Println("I'm a client.")
	return nil
}

func (p ServerProtocol) Run() error {
	fmt.Println("I'm a server.")
	return nil
}

func main() {
	client := ClientProtocol{Id: 76}
	server := ServerProtocol{}

	protocols := [2]Protocol{client, server}
	protocols[0].Run()
	protocols[1].Run()
}
```

### Type Assertion
Type assertions are similar to dynamic type casting in Java and C#. The notation `x.(T)` is called type assertion. In the Protocol example above, we can write:

```go
cp, ok := protocols[0].(ClientProtocol)		// Type assertion
if ok {
	fmt.Println(cp.Id)
} else {
	fmt.Println("Not a client protocol!")
}
```