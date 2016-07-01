package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Message struct {
	Name string
	Body string
	Time int64
}

type IntStringMap map[int]string

func (iMap IntStringMap) Marshal() ([]byte, error) {
	x := make(map[string]string)
	for k, v := range iMap {
		x[strconv.Itoa(k)] = v
	}
	return json.Marshal(x)
}

func (iMap *IntStringMap) Unmarshal(b []byte) error {
	x := make(map[string]string)
	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}
	*iMap = make(IntStringMap, len(x))
	for k, v := range x {
		ki, _ := strconv.Atoi(k)
		(*iMap)[ki] = v
	}
	return nil
}

func main() {
	// Marshaling/unmarshaling a struct
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, err := json.Marshal(m) // Note: json package only accesses exported fields of struct types
	if err == nil {
		var x Message
		err = json.Unmarshal(b, &x)
		if err == nil {
			fmt.Println("Unmarshaled struct: ", x)
		}
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Marshaling/unmarshaling a map
	s := make(map[string]int) // Only map[string] T can be marshaled by the json package
	s["hello"] = 123456
	s["bye"] = 78901
	b, err = json.Marshal(s)
	if err == nil {
		var x map[string]int
		err = json.Unmarshal(b, &x)
		if err == nil {
			fmt.Println("Unmarshaled map: ", x)
		}
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Marshaling/unmarshaling a matrix
	a := make([][]int, 3)
	a[0] = []int{1, 2, 3}
	a[1] = []int{4, 5, 6}
	a[2] = []int{7, 8, 9}
	b, err = json.Marshal(a)
	if err == nil {
		var x [][]int
		err = json.Unmarshal(b, &x)
		if err == nil {
			fmt.Println("Unmarshaled matrix: ", x)
		}
	} else {
		fmt.Println("Error:", err.Error())
	}

	// Marshalling/unmarshalling an int->string map which is not directly possible with JSON
	p := make(IntStringMap)
	p[123] = "dragon"
	p[4567] = "fly"

	b, err = p.Marshal()
	if err == nil {
		var y IntStringMap
		err = y.Unmarshal(b)
		if err == nil {
			fmt.Println("Unmarshaled map: ", y)
		}
	} else {
		fmt.Println("Error:", err.Error())
	}
}
