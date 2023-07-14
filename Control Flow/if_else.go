package main

import "fmt"

func main() {
	a := "Zello"
	b := "Trell"

	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a")
	}
}

// you can only put 'else' or 'else if' at the end of the if block now below the block end
