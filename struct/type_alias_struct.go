package main

import "fmt"

type Person struct {
	name string
	age  int
}

type People1 struct {
	name string
	age  int
}

type People2 = struct {
	name string
	age  int
}

type People = Person

func main() {
	var p3 = People2{"bod", 21}
	var p1 = Person{"wick", 21}
	fmt.Println(p1)
	fmt.Println(p3)

	var p2 = People{"coba yang lain", 23}
	fmt.Println(p2)
}
