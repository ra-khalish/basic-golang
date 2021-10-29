package main

import "fmt"

func main(){
	// menggunakan var, menggunakan tipe data, menggunakan perantara "="
	var firstName string = "rafi"

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName := "khalish"

	fmt.Printf("halo %s %s!\n", firstName, lastName)

	// assignment variabel
	hisName := "lu"
	hisName = "lul"
	hisName = "lulu"
	fmt.Printf("halo %s!\n", hisName)
}
