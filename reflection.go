package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Age    int
	Name   string
	Height float64
}

func main() {
	var x int = 56
	fmt.Println("type:", reflect.TypeOf(x))

	t := Person{26, "Mahdi", 5.7}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%s %s = %v\n", typeOfT.Field(i).Name, f.Type(), f.Interface())
	}
}
