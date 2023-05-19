package main

import "fmt"

type I interface {
	M()
}

// precisaria de alguma implementação de I
// para passar pelo menos o tipo que a implementa  para a variável 'i'
// assim como está, o codigo dá erro de runtime
func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
