package main

import "fmt"

func val_seman(a int) int {
	a = 10
	return a

}

func poi_seman(a *int) int {

	*a = 10
	return *a
}
func main() {

	e := 5

	fmt.Println(val_seman(e))
	fmt.Println(e)
	fmt.Println(poi_seman(&e))
	fmt.Println(e)

}
