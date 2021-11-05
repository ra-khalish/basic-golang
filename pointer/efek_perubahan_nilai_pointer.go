package main

import "fmt"

func main() {
	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value) :", numberA) // 4
	fmt.Println("numberA (address) :", &numberA) // referencing dengan &

	fmt.Println("numberB (value) :", *numberB) // 4 deferencing dengan *
	fmt.Println("numberB (address) :", numberB) // referencing

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value) :", numberA) // 4
	fmt.Println("numberA (address) :", &numberA) // referencing dengan &

	fmt.Println("numberB (value) :", *numberB) // 4 deferencing dengan *
	fmt.Println("numberB (address) :", numberB) // referencing

}
