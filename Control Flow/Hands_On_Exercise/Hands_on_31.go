package main

import "fmt"

func main() {
	mp := map[string]int{
		"Subrat":  1,
		"Shubham": 2,
		"Khush":   3,
		"Shanku":  4,
	}

	for key, value := range mp {
		fmt.Printf("%v  %v \n", key, value)
	}
}
