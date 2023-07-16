package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("The value x and y is %v and %v respectively \n", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Println("x and y are less than 4")
	case x > 6 && y > 6:
		fmt.Println("x greater than 6 and y greater than 6")
	case x >= 4 && x <= 6:
		fmt.Println("x greater or equal to 4 and less than equal to 6")
	case y != 5:
		fmt.Println("y not equal to 5")
	default:
		fmt.Println("none of the cases were met")
	}
}
