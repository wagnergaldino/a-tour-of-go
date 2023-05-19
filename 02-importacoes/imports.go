package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randnumber := rand.Float64()
	fmt.Printf("randnumber = %v \nNow you have %g problems.\n", randnumber, math.Sqrt(randnumber))
}
