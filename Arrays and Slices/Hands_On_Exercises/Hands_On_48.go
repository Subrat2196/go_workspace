package main

import "fmt"

func main() {
	xi := []string{"James", "Bond", "Shaken, not stirred"}
	xj := []string{"Miss", "MoneyPenny", "I'm 008"}

	xxi := [][]string{xi, xj}

	fmt.Println(xxi)

	for i, v := range xxi {
		fmt.Println(i, " ", v)
		for idx, val := range v {
			fmt.Println(idx, " ", val)
		}

	}
}
