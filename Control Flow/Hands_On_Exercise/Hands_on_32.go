package main

import "fmt"

func main() {
	mp := map[string]int{
		"sub": 1,
		"dub": 2,
	}
	x := mp["dub"]
	fmt.Println(x)
	if y, ok := mp["ss"]; !ok {
		fmt.Println("There is no y")
	} else {
		fmt.Println("there is y", y)
	}
}
