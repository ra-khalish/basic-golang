package main

import (
	"belajar-golang-level-akses/library"
	"fmt"
	// contoh alias pada package
)

func main() {
	// var s1 = library.Student{"ethan", 21}
	// f.Println("name", s1.Name)
	// f.Println("grade", s1.Grade)
	sayHello("ethan")
	fmt.Printf("name : %s\n", library.Student.Name)
	fmt.Printf("grade : %d\n", library.Student.Grade)
}
