package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	var s1 = struct {
		person
		grade int
	}{} // bisa di isi dengan value ke obj
	s1.person = person{"wick", 21}
	s1.grade = 2

	fmt.Println("name :", s1.person.name)
	fmt.Println("age :", s1.person.age)
	fmt.Println("grade :", s1.grade)
}
