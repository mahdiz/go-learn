package main

import "fmt"

type Vertex struct {
	x int
	y int
}

func (v Vertex) addVertex1(v2 Vertex) {
	v.x += v2.x
	v.y += v2.y
}

func (v *Vertex) addVertex2(v2 Vertex) {
	v.x += v2.x
	v.y += v2.y
}

func main() {
	v1 := Vertex{x: 1, y: 2}
	v2 := Vertex{x: 3, y: 4}
	v1.addVertex1(v2)
	fmt.Println(v1)
	v1.addVertex2(v2)
	fmt.Println(v1)
}
