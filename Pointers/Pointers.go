package main

import "fmt"

func main() {

	x := 3
	// & is the referencing/address operator
	// * is the dereferencing operator
	fmt.Println(x, &x)

	fmt.Printf("%v\t%T\n", &x, &x)

	y := &x

	fmt.Println(x, *y, &x, &y)

}
