package main

import "fmt"

// iota is used to initialize constant variables in a sequengtial order
// iota always start with 0 for numeric integer constants
func main() {
	const (
		_ = iota
		jan
		feb
		march
		april
		may
	)
	fmt.Println(jan, feb, march, april, may)

}
