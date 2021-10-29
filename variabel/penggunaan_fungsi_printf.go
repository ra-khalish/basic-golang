package main

import "fmt"

func main(){
	var firstName string = "rafi"

	var lastName string
	lastName = "khalish"

	fmt.Printf("halo rafi khalish!\n")
	fmt.Printf("halo %s %s!\n", firstName, lastName)
	fmt.Println("halo", firstName, lastName + "!")
}
