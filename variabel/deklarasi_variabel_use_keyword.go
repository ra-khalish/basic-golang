package main

import "fmt"

func main(){
	name := new(string)

	fmt.Println(name) // nilai alamat memori
	fmt.Println(*name) // perlu di deference dengan "*"
}
