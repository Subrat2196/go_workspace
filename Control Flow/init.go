package main

import "fmt"

func init() {
	fmt.Println("Yes said the init function,oops i am already above where u asked")
}

func main() {
	fmt.Println("Can you run before me said the main function")
}
