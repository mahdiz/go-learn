package main

import "fmt"

type IProtocol interface {
	Run() error
}

type Protocol struct {
	Id int
}

type ClientProtocol struct {
	Protocol
	Username string
}

type ServerProtocol struct {
	Protocol
	IsUp bool
}

func (p *ClientProtocol) Run() error {
	fmt.Println("I'm a client with id =", p.Id)
	return nil
}

func (p *ServerProtocol) Run() error {
	fmt.Println("I'm a server with id =", p.Id)
	return nil
}

func main() {
	client := ClientProtocol{Username: "mahdiz", Protocol: Protocol{Id: 76}}
	server := ServerProtocol{IsUp: true, Protocol: Protocol{Id: 248}}

	protocols := [2]IProtocol{&client, &server}
	protocols[0].Run()
	protocols[1].Run()

	cp, ok := protocols[0].(*ClientProtocol) // Type assertion
	if ok {
		fmt.Println(cp.Username)
		fmt.Println(cp.Id)
	} else {
		fmt.Println("Not a client protocol!")
	}

	var anything [2]interface{}
	anything[0] = client
	anything[1] = server
	ca, ok := anything[0].(ClientProtocol) // Type assertion
	if ok {
		fmt.Println(ca.Username)
		fmt.Println(ca.Id)
	} else {
		fmt.Println("Not a client protocol!")
	}
}
