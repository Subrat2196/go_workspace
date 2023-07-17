package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(xi)

	// deleting using slicing a slice and append methods

	// xj := xi[5:]
	// xk := xi[:4]

	// xl := append(xk, xj...)
	// fmt.Println(xl)

	// rather than the above step we can directly do using append function

	xi = append(xi[:4], xi[5:]...)
	fmt.Println(xi)

}
