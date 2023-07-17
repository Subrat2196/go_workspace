package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6}
	xj := []int{7, 8, 9}
	xxi := [][]int{xi, xj}

	fmt.Println(xi)
	fmt.Println(xj)
	fmt.Println(xxi)

}

// multi dimensional slices means slice of slices
// e.g a := []string {"a","b","c"}  ,we can read it as each element in slice named a is string
// Now we have e.g   a := [][]string{slice 1,slice 2} , here each element in slice is []string i.e slice
