package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

func main() {

	p1 := person{
		first:  "Subrat",
		second: "Bahuguna",
		age:    27,
	}

	fmt.Println(p1)

	mp := map[string]person{
		p1.second : p1
	}

	

}
