package main

import "fmt"

func main(){
	var data map[string]int
	// data["one"] = 1
	// muncul error?
	// fmt.Println(data)

	data = map[string]int{}
	data["one"] = 1
	fmt.Println(data)
	// tidak ada error

	// cara horizontal
	var chicken1 = map[string]int{"januari": 50, "februari": 40}

	//cara vertical
	var chicken2 = map[string]int{
		"januari": 50,
		"februari": 40,
	}

	var chicken3 = map[string]int{}
	var chicken4 = make(map[string]int)
	var chicken5 = *new(map[string]int)

	fmt.Println(chicken1)
	fmt.Println(chicken2)
	fmt.Println(chicken3)
	fmt.Println(chicken4)
	fmt.Println(chicken5)
}
