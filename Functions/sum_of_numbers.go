package main

import "fmt"

func main() {

	addition := add(1, 3, 4, 5, 6, 7, 7, 10, 12, 124)

	fmt.Println(addition)

}

func add(ii ...int) int {
	sum := 0
	for _, val := range ii {
		sum += val
	}

	return sum
}
