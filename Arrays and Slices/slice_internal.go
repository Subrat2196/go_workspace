package main

import "fmt"

func main() {
	a := []int{4, 3, 2, 1, 5, 6, 7, 8}

	// b := a
	// here what happens is since now b is pointing to the same address as a
	// any change in a or b will impact the other also

	b := make([]int, 8)
	fmt.Println(b)
	// using the copy function here we are copying the a elements to b
	copy(b, a)
	// now since b is an entirely different slice as we created from make function ,
	// a and b are not related
	fmt.Println(b)

	// here if we want to pass an array to a function
	// calling - to_sort(a)
	// function definition - func to_sort(x []int) int {}
	// here we can see that now x will take value of a so it will be like x := a
	// therefore any change in 'x' will cause change in 'a'
	// solution is to first copy(b,x) then work on b instead

}
