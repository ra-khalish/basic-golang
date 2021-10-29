package main

import "fmt"

func main(){
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"

	var fourth, fifth, sixth string = "empat", "lima", "enam"

	// lebih ringkas
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"

	// deklarasi multi variabel dengan teknik inference
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Println(first, second, third, fourth, fifth, sixth, seventh, eight, ninth)
	fmt.Println(one, isFriday, twoPointTwo, say)
}
