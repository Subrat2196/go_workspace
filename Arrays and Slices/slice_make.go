package main

import (
	"fmt"
)

func main() {
	xi := make([]int, 0, 10) // make (type, len, capacity) , here capacity is compulsory to give
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))

	xi = append(xi, 1, 2, 3, 4, 5, 6, 7, 8)
	xi = append(xi, 9, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println(len(xi))
	fmt.Println(cap(xi))
}

// Length is the number of elements at that particular time in an array/slice
// capacity is the maximum number of elements that can be stored
// we can use append to add elements also

//capacity becomes double whenever the existing capacity is overflowed
