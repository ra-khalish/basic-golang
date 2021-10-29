package main

import "fmt"

func main()  {
	var numbers = []int{2, 3, 0, 4, 3, 2, 0, 4, 2, 0, 3}
	params := 3
	var newNumbers = func(min int) []int {
		var r []int
		for _, e := range numbers {
			if e < min {
				continue
			}
			r = append(r, e)
		}
		return r
	}(params)// Ciri khas IIFE adanya kurung paramter tepat setelah deklarasi
	// Closure berakhir. Jika ada parameter, bisa juga dituliskan dalam kurung
	// Parameternya
	fmt.Println("original number :", numbers)
	fmt.Println("filtered number :", newNumbers)
}