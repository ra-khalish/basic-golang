package main

import "fmt"

func main(){
	const firstName string = "rafi"
	fmt.Print("halo ", firstName, "!\n")

	const lastName = "khalish"
	fmt.Print("nice to meet you ", lastName, "!\n")

	/*
	masukan nilai baru ke konstanta lastName
	error output "cannot assign to lastName"
	
	lastName = "khalish baru"
	fmt.Print(lastName)
	*/

	// Penggunaan fungsi fmt.Print()
	fmt.Println("Contoh output fungsi Println():")
	fmt.Println("rafi khalish")
	fmt.Println("john", "wick")

	// Print() tidak membuat sepasi otomatis
	fmt.Println("Contoh output fungsi Print():")
	fmt.Print("rafi khalish\n")
	fmt.Print("rafi ", "khalish\n")
	fmt.Print("rafi", " ", "khalish\n")
}
