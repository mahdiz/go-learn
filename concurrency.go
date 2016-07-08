package main

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func Channels() {
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

// Single goroutine
func Concurrency1() {
	channel := make(chan bool)

	go func() {
		fmt.Println("Waiting two seconds...")
		time.Sleep(2 * time.Second)
		fmt.Println("Time elapsed.")
		channel <- true
	}()

	fmt.Println("Goroutine started...")

	// Wait until something is written on the channel
	<-channel

	fmt.Println("Goroutine finished!")
}

// Multiple goroutine
func Concurrency2() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		delay := rand.Intn(5)
		time.Sleep(time.Duration(delay) * time.Second)
		c1 <- delay
	}()

	go func() {
		delay := rand.Intn(5)
		time.Sleep(time.Duration(delay) * time.Second)
		c2 <- delay
	}()

	fmt.Println("Goroutines started...")

	// Wait until something is written on both channels
	for i := 0; i < 1; i++ {
		select {
		case delay := <-c1:
			fmt.Println("First goroutine finished after " + strconv.Itoa(delay) + " sec.")
		case delay := <-c2:
			fmt.Println("Second goroutine finished after " + strconv.Itoa(delay) + " sec.")
		}
	}
	fmt.Println("We're done.")
}

func Concurrency3() {
	ch := make(chan int)
	go func() {
		for {
			fmt.Println("Writing sth on the channel")
			ch <- 56
			time.Sleep(1000 * time.Millisecond)
		}
	}()

	for {
		select {
		case newCh := <-ch:
			fmt.Println(newCh)
		}
	}
}

func main() {
	//Concurrency3()

	s := "192.168.0.1:9572"

	//u := url.Parse(s)
	host, port, _ := net.SplitHostPort(s)
	fmt.Println(host, port)
}
