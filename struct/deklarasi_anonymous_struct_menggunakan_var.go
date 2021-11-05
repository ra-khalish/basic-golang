package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	// hanya deklarasi
	var student struct {
		person
		grade int
	}
	// deklarasi sekaligus inisialisasi
	var student1 = struct {
		grade int
	}{12}

	student.person = person{"wick", 21}
	student.grade = 2
	student1.grade = 3
	fmt.Println("person", student.person)
	fmt.Println("student grade ", student.grade)
	fmt.Println("student1 grade ", student1.grade)
}
