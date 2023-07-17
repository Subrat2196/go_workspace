// a numbered sequence of elements of the same type
// fixed size
// used for Go Internals; generally not used in your code
// we mostly use slice because they are dynamic and can be changed
// slices have length and capacity
package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(a)

	b := [5]int{1, 2, 3, 3, 3}
	fmt.Println(b)

	b = a
	fmt.Println(b)

	fmt.Println(len(b))
}

// arrays have a block level scope only i.e you can only access elements if they are
// in the same block as array definition
// one arrays can be assigned to other one only if length and type of arrays are same.
