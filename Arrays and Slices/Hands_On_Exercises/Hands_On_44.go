package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xi)

	xj := xi[:5]
	xk := xi[5:]
	xl := xi[2:7]
	xm := xi[1:6]

	fmt.Println(xj, xk, xl, xm)
}
