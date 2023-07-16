package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(250)

	fmt.Printf("The name of the variable is x and the value is %v", x)

	if x <= 100 {
		for i := 1; i < 100; i++ {
			fmt.Print(" ", i)
		}
	} else if x > 100 && x <= 200 {
		for i := 100; i < 200; i++ {
			fmt.Print(" ", i)
		}
	} else {
		for i := 200; i < 250; i++ {
			fmt.Print(" ", i)
		}
	}

}
