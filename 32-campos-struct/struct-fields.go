package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println("Vertex{x,y} =", v)
	v.X = 4
	fmt.Println("x do Vertex mudado para", v.X)
	fmt.Println("Vertex{x,y} =", v)
}
