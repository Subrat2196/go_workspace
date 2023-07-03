package main

import "fmt"

func swap(a, b string) (string, string) {
	return b, a
}

func main() {
	// The swap function return two string

	s1 := "string 1"
	s2 := "string 2"
	fmt.Println(s1, s2)
	s3, s4 := swap(s1, s2)
	fmt.Println(s3, s4)

}
