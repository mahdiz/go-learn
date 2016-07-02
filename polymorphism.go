package main

import "fmt"

type ClientProtocol struct {
	Id int
}

type ServerProtocol struct {
	IsUp bool
}

type Protocol interface {
	Run() error
}

type Anything interface{}

func (p *ClientProtocol) Run() error {
	fmt.Println("I'm a client.")
	return nil
}

func (p *ServerProtocol) Run() error {
	fmt.Println("I'm a server.")
	return nil
}

func main() {
	client := ClientProtocol{Id: 76}
	server := ServerProtocol{}

	protocols := [2]Protocol{&client, &server}
	protocols[0].Run()
	protocols[1].Run()

	cp, ok := protocols[0].(*ClientProtocol) // Type assertion
	if ok {
		fmt.Println(cp.Id)
	} else {
		fmt.Println("Not a client protocol!")
	}

	anything := [2]Anything{&client, &server}
	ca, ok := anything[0].(*ClientProtocol) // Type assertion
	if ok {
		fmt.Println(ca.Id)
	} else {
		fmt.Println("Not a client protocol!")
	}
}
