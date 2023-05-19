package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	// Gerará panico pq i é tipo string e a atribuição não está sendo verificada
	// f = i.(float64) // panic
	// fmt.Println(f)
}
