package main

import "fmt"

func main() {

	mp := map[string]int{
		"Subrat": 1,
		"Kishen": 2,
		"Suhail": 3,
	}

	fmt.Println(mp)

	// we have a built in keyword for map which is delete

	delete(mp, "Kishen")
	fmt.Println(mp)
}

// even if we search for a key that doesn't exist or find the value of a key
// that doesn't exist it will give 0
// even if you delete from a map a key that don't exist it won't panic
