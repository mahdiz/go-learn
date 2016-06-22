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

## Object-Oriented
A struct is a collection of fields: 

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

Instead, we can use `interface`:

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