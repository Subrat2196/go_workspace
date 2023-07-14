package main

import "fmt"

func main() {
	var name string = "Kim"
	var age int = 20

	fmt.Printf("The age of %s is %d and the type of them is %T and %T \n ", name, age, name, age)
	// %s and %d are known as format verbs or format specifiers for string and integer
	// we can use %v whenever we want to take value
	fmt.Printf("The age of %v is %v and the type of them is %T and %T \n", name, age, name, age)
}
