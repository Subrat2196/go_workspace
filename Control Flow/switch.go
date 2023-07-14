package main

import "fmt"

func main() {

	a := 47

	switch {
	case a == 42:
		fmt.Println("a is 42")
		fallthrough
	case a > 42:
		fmt.Println("a is greater than 42")
		fallthrough
	case a < 42:
		fmt.Println("a is smaller than 42")
	}
}

// Default is not compulsory in Switch statement
//
