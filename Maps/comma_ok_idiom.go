// We know that if we want to know a value of a non existent key it will give 0
// to find a key exist or not we can use comma_ok_idiom

package main

import "fmt"

func main() {
	mp := map[string]int{
		"Daisy": 1,
		"leona": 2,
	}

	fmt.Println(mp)
	// comma ok is used here to see if present or not present
	v, ok := mp["dassi"]
	fmt.Println(v, ok)

}
