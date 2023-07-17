package main

import "fmt"

func main() {
	xi := []string{"Hello", "I", "Am", "Yours"}
	xj := []string{"You", "Are", "Mine"}

	xk := append(xi, "And", "You", "know", "what")
	xl := append(xi, xj...)

	fmt.Println(xk)
	fmt.Println(xl)
}
