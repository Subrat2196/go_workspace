package main

import (
	"fmt"
)

func add(x, y int) int {
	return x + y
}

func main() {

	a, b := 2, 3
	c := add(a, b)
	fmt.Printf("The addition of two numbers is %v \n", c)

}
