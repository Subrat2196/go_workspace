package main

import "fmt"

func add(x *int, y *int) int {
	z := *x + *y
	return z
}

func main() {
	a := 10
	b := 20

	z := add(&a, &b)

	fmt.Println(z)

}

// here in go many data structures like array, structures, maps,slices,
// channels,interfaces, functions, etc. are all reference types.
