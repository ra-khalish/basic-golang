package main

import "fmt"

func main(){
	var fruits = [4]string{"apple", "grape", "banana", "melon"}

	// jika menampilkan nilai pada elemen array
	for _, fruit := range fruits {
		fmt.Printf("nama buah: %s\n", fruit)
	}
	// jika hanya ingin menampilkan indeks elemen array
	/*
	for i, _ := range fruits { }
	// atau
	for i := range fruits { }
	*/
}
