package main

import "fmt"

func main(){
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0]) // "apple"

	/* jika jumlah elemen tidak dituliskan,
	maka variabel tersebut adalah slice.
	*/
	var fruitsA = []string{"apple", "grape"} // slice
	var fruitsB = [2]string{"banana", "melon"} // array
	var fruitsC = [...]string{"papaya", "grape"} // array

	fmt.Println(fruitsA)
	fmt.Println(fruitsB)
	fmt.Println(fruitsC)

}
