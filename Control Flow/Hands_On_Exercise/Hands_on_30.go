// here we are going to create a slice and we will iterate through it

package main

import "fmt"

func main() {
	xi := []int{42, 43, 44, 45, 46}

	for i, v := range xi {
		fmt.Printf("The index is %v and value is %v\n", i, v)
	}

}
