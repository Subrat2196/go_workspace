package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46, 47, 48}

	fmt.Println(xi)

	for v, t := range xi {
		fmt.Printf("%v - %T \n", v, t)
	}
}
