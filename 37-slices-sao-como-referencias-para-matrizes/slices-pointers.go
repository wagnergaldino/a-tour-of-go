package main

import "fmt"

func main() {
	// matriz
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	// slices
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// mudando um valor no slice muda tb a matriz subjacente
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
