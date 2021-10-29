package main

import "fmt"

func main(){
	var names [4]string
	names[0] = "rafi"
	names[1] = "dan"
	names[2] = "lulu"
	names[3] = "!"

	fmt.Println(names[0], names[1], names[2], names[3], names[4])
}
