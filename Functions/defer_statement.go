// What defer statement say is don't run this function call until the function this call is enclosed
// in ends or finishes

package main

import "fmt"

func add(ii ...int) {
	sum := 0
	for _, val := range ii {
		sum += val
	}

	fmt.Printf("The sum is %v and i will run first as i am above you in execution \n", sum)

}

func sayit(strr string) {
	fmt.Println(strr)
}

func main() {
	// here we will first create a variadic function

	s := []int{1, 2, 3, 3, 4, 5, 6}
	defer add(s...)

	str := "Hey i'll put you in defer statement and i'll run first ha.ha..hahaha"
	sayit(str)

}
