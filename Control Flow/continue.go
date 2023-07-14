package main

import "fmt"

func main() {
	// this is the program to print even number till a particular number using for loop and continue

	till := 20

	for i := 0; i <= till; i++ {
		if i%2 == 0 {
			fmt.Print(" ", i)
		}
	}

}
