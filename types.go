package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

type Employee struct {
	Person // Called embedding (not inheritance)
	Salary int
	Bonus  int
}

// Some method for Employee
// A method is a function declared with a receiver. A receiver is a value or a pointer of a struct type.
func (e *Employee) TotalSalary() int {
	return e.Salary + e.Bonus
}

type Computer struct {
	IpAddress string
	Model     string
}

// An interface to allow upcasting (OOP speaking)
// Interfaces specify that objects of a particular types can behave in a specific way
// Main flexibility of interfaces Any type that satisfies the definition of the interface
// is implicitly bound to the interface although those types do not explicitly mention this.
type Entity interface {
	// Entity is anything that can provide some ID
	Id() string
}

// Implement interface methods for previously-defined types and those types "implicitly" bind to the interface
func (p Person) Id() string {
	// Note that the method doesn't have pointer receiver (p *Person)
	return p.FirstName + " " + p.LastName
}

func (c Computer) Id() string {
	return c.IpAddress
}

func TestTypes() {
	e := Employee{Person: Person{FirstName:"Mahdi", LastName:"Zamani"}, Salary:75000, Bonus:10000}
	c := Computer{IpAddress:"203.162.188.151", Model:"iMac"}

	entities := [...]Entity{e, c}

	for _, entity := range entities {
		fmt.Println(entity.Id())
	}
}