//here we will learn the creation,accessing of elements,printing our map,length of map

package main

import "fmt"

func main() {

	mp := map[string]int{
		"subrat":   1,
		"bahuguna": 2,
	}
	fmt.Println(mp)

	mp["subrat"] = 21

	fmt.Println(mp)

	// now we will create a map using make keyword

	mp2 := make(map[string]int)

	mp2["Hello"] = 1
	mp2["Haallo"] = 2
	mp2["Ma'am"] = 3

	fmt.Println(mp2, len(mp2))

	for key, val := range mp2 {
		fmt.Println(key, val)
	}

	for _, val := range mp2 {
		fmt.Println(val)
	}

}
