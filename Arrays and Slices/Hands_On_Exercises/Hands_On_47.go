package main

import "fmt"

func main() {
	// states := []string{"Washington", "Cleveland", "orlando", "San Francisco"}
	states := make([]string, 0, 50)
	states = append(states, "Colarado", "Washington", "San_Francisco", "Orlando")
	fmt.Printf("The length,capacity and values are %v %v %v \n", len(states), cap(states), states)

}
