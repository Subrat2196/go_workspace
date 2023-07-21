// Unfurling a slice means that when we are using a variadic parameter than we need something
// which can take any number of values e.g func add(ii ...int), here when we call it like :-
// add(1,2,3,4,5,6,7,8,9) , then it will take all the values from 1 - 9 one by one.
// now what it if we want to give all the values that a slice contains and use them to give values
// to a function, this is called unfurling a slice, i.e we open a slice and pass it's values
// as parameters for a function

// e.g we have a slice s that is initialized , then we can pass it like add(s...)
// three dots after the slice name is how we pass the parameter
package main

import "fmt"

func add(ii ...int) int {
	sum := 0
	for _, val := range ii {
		sum += val
	}

	return sum
}

func main() {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8}

	sum := add(s...)

	fmt.Println(sum)
}
