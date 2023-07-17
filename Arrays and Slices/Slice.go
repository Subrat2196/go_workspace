package main

import "fmt"

func main() {

	xi := []int{3, 4, 5, 6, 7, 8}

	for idx, val := range xi {
		fmt.Println(idx, val)
	}

	// fmt.Println(xi) can also be used to print array

	// printing only values or only index we need to give blank identifier

	for _, val := range xi {
		fmt.Println(val)
	}

}
