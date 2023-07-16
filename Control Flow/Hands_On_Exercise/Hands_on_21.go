package main

import (
	"fmt"
	"math/rand"
)

// Niladic function means a function that takes no arguments e.g init function
// doesn't take any arguments

func init() {
	fmt.Println("This is where the initialization of the program begins")
}

func main() {
	x := rand.Intn(250)
	fmt.Printf("%v\n", x)

	switch {
	case x <= 100:
		for i := 0; i <= 100; i++ {
			fmt.Print(" ", i)
		}
	case x > 100 && x <= 200:
		for i := 100; i <= 200; i++ {
			fmt.Print(" ", i)
		}
	case x > 200:
		for i := 200; i <= 250; i++ {
			fmt.Print(" ", i)
		}
	}

}
