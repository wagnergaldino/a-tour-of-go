package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(x, y int) int {
	return x + y
}

func main() {
	rand.Seed(time.Now().UnixNano())
	rn1 := rand.Intn(100)
	rn2 := rand.Intn(100)
	fmt.Printf("add(%v, %v) = %v", rn1, rn2, add(rn1, rn2))
}
