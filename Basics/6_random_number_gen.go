package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	a := 10
	fmt.Println(rand.Intn(a), math.Sqrt(34))
}

// rand.Intn(number) -> generates a random number from [0,number)
