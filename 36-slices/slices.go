package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("primes =", primes)

	//inclui o primeiro e exclui o ultimo
	var s []int = primes[1:4]
	fmt.Println("primes[1:4] =", s)
}
