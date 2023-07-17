package main

import "fmt"

// we can create our own type

type person struct {
	first_name string
	last_name  string
	age        int
}

func main() {

	p1 := person{
		first_name: "Subrat",
		last_name:  "Bahuguna",
		age:        24,
	}

	p2 := person{
		first_name: "Khushboo",
		last_name:  "Sahu",
		age:        10,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first_name)

}
