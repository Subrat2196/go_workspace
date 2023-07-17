package main

import "fmt"

func main() {
	// notation for slicing is [inclusive : exclusive]
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(xi)
	// here we are slicing a slice from index [1,3)
	xj := xi[1:3]
	fmt.Println(xj)

	xk := xi[:7]
	fmt.Println(xk)

	xl := xi[:]
	fmt.Println(xl)

}
