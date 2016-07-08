package main

import "fmt"

const (
	PROTOCOL_TYPE_CLIENT = iota
	PROTOCOL_TYPE_SERVER
)

type IProtocol interface {
	Type() int
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

func (p *ClientProtocol) Type() int {
	return PROTOCOL_TYPE_CLIENT
}

func (p *ClientProtocol) Run() error {
	fmt.Println("I'm a client with id =", p.Id)
	return nil
}

func (p *ServerProtocol) Type() int {
	return PROTOCOL_TYPE_SERVER
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

	for _, p := range protocols {
		switch p.Type() {
		case PROTOCOL_TYPE_CLIENT:
			cp, _ := p.(*ClientProtocol) // Type assertion
			fmt.Println("Client username:", cp.Username)
			fmt.Println("Client id:", cp.Id)

		case PROTOCOL_TYPE_SERVER:
			sp, _ := p.(*ServerProtocol) // Type assertion
			fmt.Println("Server is up?", sp.IsUp)
			fmt.Println("Server id:", sp.Id)
		}
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
