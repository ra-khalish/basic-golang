package main

import "fmt"

type student struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

func main() {
	var s = student{
		person: struct {
			name string
			age  int
		}{},
		grade:   0,
		hobbies: []string{},
	}
	fmt.Println(s)

}
