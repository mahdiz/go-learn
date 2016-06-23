// Basic Go types:
// bool, byte, string
// int, int8, int16, int32, int64
// uint, uint8, uint16, uint32, uint64
// float32, float64

package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os/user"
	"time"
)

func basics() {

	// Reading an integer from stdin:
	var x int
	fmt.Scanf("%v\n", &x)

	// Reading an integer and a char right after it:
	var c byte
	fmt.Scanf("%v%c\n", &x, &c)
	fmt.Println(string(c))

	// Writing to stdout
	fmt.Println("Test")
	fmt.Print("Test\n")
	fmt.Printf("%v %s", x, "Test")
	fmt.Printf("%b", x) // Print base 2 of integer x

	// Outputting decimal numbers
	y := 87.567013
	fmt.Printf("%.3f", y) // Prints 87.567
	z := 7
	fmt.Printf("%03d", z) // Prints 007

	// Casting to another type
	n := 26
	s := float32(n)
	fmt.Println(s)

	// Loop
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	// Random number generation
	x = rand.Intn(1000) // A random number between [0,1000)
	fmt.Println()
}

func mathExample() {

	x := 7
	y := 3
	z := math.Ceil(float64(x) / float64(y))

	fmt.Println(z)
}

func arrayExample() {

	arr1 := [3]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6} // unknown size notation

	// the following copies arr2 to arr1 element by element
	arr1 = arr2
	arr2[2] = 23
	fmt.Println(arr1, arr2)
}

func sliceExample() {

	arr := [5]int{1, 2, 3, 4, 5}
	slc1 := arr[1:4]
	slc2 := arr[3:]
	slc3 := arr[:4]

	slc4 := append(slc2, 99)
	fmt.Println(slc1, slc2, slc3, slc4)

	// len(arr1) gives the length of the array while cap(arr1) gives the the block of memory reserved for the array.
	fmt.Println(len(slc4), cap(slc4))
}

// a map is a dictionary data structure
func mapExample() {

	// define the map
	var myMap map[string]int

	// initialize the map
	myMap = map[string]int{}

	// add key-value pairs
	myMap["a"] = 1
	myMap["b"] = 2

	// non-existing values will return the default value which is 0 in our case
	fmt.Println(myMap["c"])

	// to check the existance use the second return value
	val, exists := myMap["c"]
	if exists {
		fmt.Println(val)
	}

	// to delete a pair from the map
	delete(myMap, "b")
}

func funcHandler() {

	type handler func(a []string)
	var h handler
	h = func(a []string) {
		for sig := range a {
			fmt.Println(sig)
		}
	}

	h([]string{"a", "b", "c"})
}

func checkOS() {

	usr, err := user.Current()
	if err != nil {
		fmt.Println("Error")
	}
	homedir := usr.HomeDir
	fmt.Println(homedir)
}

func goroutine() {

	go func() {
		panic("Signal: ")
	}()
}

func channels() {

	a := []int{1, 2, 3, 4, 5, 6}

	// create a channel of integers
	channel := make(chan int)

	// create a goroutine to find sum and write the result to the channel
	go func(arr []int, c chan int) {
		s := 0
		for _, e := range arr {
			s += e
			time.Sleep(500 * time.Millisecond)
		}
		// write the sum to the channel
		c <- s
	}(a, channel)

	// read from the channel (will wait until something is written on the channel)
	x := <-channel
	fmt.Println(x)
}

// waits for two channels and proceeds based on which channel has something to read
func fibonacci(c, quit chan int) {

	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// uses the above fibonacci function
func channelsSelect() {

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

// timers use channels
func timer() {

	tick := time.Tick(1000 * time.Millisecond)
	boom := time.After(5000 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("Tick")
		case <-boom:
			fmt.Println("Boom!")
			return
		}
	}
}

// Marshals a list of byte arrays
// Each input array must have less than 65536 bytes (65KB)
func MarshalArrays(arrs ...[]byte) []byte {

	size := 0
	for _, arr := range arrs {
		size += len(arr) + 2
	}

	i := 0
	res := make([]byte, size)
	for _, arr := range arrs {
		len := len(arr)
		binary.BigEndian.PutUint16(res[i:i+2], uint16(len))
		copy(res[i+2:i+len+2], arr)
		i += len + 2
	}
	return res
}

// Unmarshals a list of byte arrays
func UnmarshalArrays(input []byte) [][]byte {

	var arrs [][]byte

	for i := 0; i < len(input); {
		len := int(binary.BigEndian.Uint16(input[i : i+2]))
		arr := make([]byte, len)
		copy(arr, input[i+2:i+2+len])
		arrs = append(arrs, arr)
		i += len + 2
	}
	return arrs
}

type animal struct {
}

type duck struct {
	animal
	featherCount int
}

func main() {

}
