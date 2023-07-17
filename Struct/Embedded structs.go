package main

// Embedded struct is embedding a struct inside another struct
import "fmt"

type person struct {
	first  string
	second string
	age    int
}

type secret_agent struct {
	person
	ltk bool
}

func main() {

	sa1 := secret_agent{
		person: person{
			first:  "sa1",
			second: "sa1",
			age:    37,
		},
		ltk: true,
	}

	p1 := person{
		first:  "Subrat",
		second: "Bahuguna",
		age:    25,
	}

	p2 := person{
		first:  "Shubham",
		second: "Bahuguna",
		age:    27,
	}

	fmt.Println(p1, p2, sa1)
}
