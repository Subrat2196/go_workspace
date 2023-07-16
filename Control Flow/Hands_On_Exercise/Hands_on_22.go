package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(10)
	y := rand.Intn(10)

	fmt.Printf("The value x and y is %v and %v respectively \n", x, y)

	if x < 4 && y < 4 {
		fmt.Println("x and y are less than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("x greater than 6 and y greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x greater or equal to 4 and less than equal to 6")
	} else if y != 5 {
		fmt.Println("y not equal to 5")
	} else {
		fmt.Println("none of the cases were met")
	}

}
