package main

import "fmt"

type person struct {
	name    string
	age     int
	hobbies []string
}

func main() {
	var p1 = struct {
		name string
		age  int
	}{age: 22, name: "wick"}
	var p2 = struct {
		name string
		age  int
	}{"ethan", 23}

	fmt.Println(p1)
	fmt.Println(p2)
}
