package main

import "fmt"

func main() {
	var a int = 4
	b, c, d, e, f, g := 1, 2, 3, 4, 5, "Hey there"

	const h int = 12
	// Here := is the short declaration operator
	// when the variable is initialized no need to declare the type in var, it will automatically take it

	var i, j = 3, 4

	fmt.Println(a, b, c, d, e, f, g, h, i, j)
	/* Imp : The short declaration operator only works inside a function, not on a package level scope */

	// the scope of a variable is from where it is declared till the end of program
	// No hosting of variables here

	// We use short declaration most of the time, we only use general decalration like
	// var a int only when we don't want to initialize it and give it 0 value
	//we don't use public and private here, we use the terminology "exported" and "imported"

}
