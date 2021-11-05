package main

import "fmt"

type student struct {
	name string
	grade int
}

func main() {
	var s1 = student{}
	s1.name = "wick"
	s1.grade = 2

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("student 1 :", s1.name)
	fmt.Println("student 2 :", s2.name)
	fmt.Println("student 3 :", s3.name)

	// contoh lain ngisi struct
	var s4 = student{name: "wayne", grade: 2}
	var s5 = student{grade: 2, name: "bruce"}

	fmt.Println("student 4 :", s4.name)
	fmt.Println("student 5 :", s5.name)
}
