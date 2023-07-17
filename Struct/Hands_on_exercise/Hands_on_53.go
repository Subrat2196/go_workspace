package main

import "fmt"

type person struct {
	first   string
	second  string
	flavors []string
}

func main() {

	p1 := person{
		first:   "Subrat",
		second:  "Bahuguna",
		flavors: []string{"Butter_Scotch", "Chocolate"},
	}

	p2 := person{
		first:   "Khush",
		second:  "Sahu",
		flavors: []string{"Vanilla", "rose"},
	}

	for _, val := range p1.flavors {
		fmt.Println(p1.first, " fav ice cream is", val)
	}

	for _, val := range p2.flavors {
		fmt.Println(p2.first, " fav ice cream is", val)
	}

}
