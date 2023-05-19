package main

import (
	"fmt"
	"math/rand"
	"time"
)

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randnumber := rand.Intn(100)
	fmt.Println("randnumber = ", randnumber)
	fmt.Print("split(randnumber) = ")
	fmt.Print(split(randnumber))
}
