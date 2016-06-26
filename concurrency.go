package main

import (
	"fmt"
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

func Concurrency() {
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

func main() {
	Concurrency()
}
