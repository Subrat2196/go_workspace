package main

import "fmt"

func main() {
	// lets create an array and initialise it
	a := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	xi := []int{3, 4, 5, 6}
	fmt.Printf("the type is %T\n", a)
	fmt.Printf("the type is %T\n", xi)
	// here the range goes through the index and the value

	for i, v := range xi {
		fmt.Println("Index , value", i, v)
	}

	mp := map[string]int{
		"Subrat":   25,
		"Khushboo": 24,
	}

	for k, va := range mp {
		fmt.Println("Key : Value = ", k, va)
	}

}
