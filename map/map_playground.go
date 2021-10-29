package main

import "fmt"

func main(){
	m := make(map[string]float64)

	m["pi"] = 3.14
	m["pi"] = 3.1416

	fmt.Println(m)

	v := m["pi"]
	fmt.Println(v)
	v = m["pie"]
	fmt.Println(v)

	e, found := m["pi"]
	fmt.Println(found, e)
	e, found = m["pie"]
	fmt.Println(found, e)

	if x, found := m["pi"]; found {
		fmt.Println(x)
	}

	delete(m, "pi")
	fmt.Println(m)
}
