// Functions are a way to reuse your code
// Here we tend to do structural/procedural programming rather than writing sphagetti code
// Functions make code more understandable
// modularization of code
// code abstraction

/*
syntax of function :-
what i used to think :-

	func name_of_func (parameters) return_type{
		body
	}

what it actually is :-

	func(receiver) name_of_func(parameters) return_type{
		body
	}

	we have recveiver also here

everything is pass by value in GO
*/
package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func aloha(s string) string {
	return fmt.Sprint("They call me Mr.", s)
}

func main() {

	first_num := 2
	second_sum := 3
	// Sprintf,Sprintl,Sprint returns a string by printing it as a string
	sum_of_two_num := add(first_num, second_sum)
	fmt.Println(sum_of_two_num)

	str := aloha("Subrat")

	fmt.Println(str)

}
