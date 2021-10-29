package main

import "fmt"

func main(){
	var chickens = []map[string]string{
		//map[string]string{"name": "chicken blue", "gender": "male"},
		//map[string]string{"name": "chicken red", "gender": "male"},
		//map[string]string{"name": "chicken yellow", "gender": "female"},
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "male"},
		{"name": "chicken yellow", "gender": "female"},
	}
	for i, chicken := range chickens {
		fmt.Println(i, chicken["gender"], chicken["name"])
	}

	var data = []map[string]string{
		{"name": "chicken blue", "gender": "male", "color": "brown"},
		{"address": "mangga street", "id": "k001"},
		{"community": "chicken lovers"},
	}
	fmt.Println(data)
}
