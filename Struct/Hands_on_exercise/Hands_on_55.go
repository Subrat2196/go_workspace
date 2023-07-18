package main

import "fmt"

type engine struct {
	electric bool
}
type car struct {
	engine
	make  bool
	model int
	doors int
	color string
}

func main() {
	c1 := car{
		engine: engine{
			electric: true,
		},
		make:  true,
		model: 350,
		doors: 4,
		color: "Red",
	}

	c2 := car{
		engine: engine{
			electric: false,
		},
		make:  true,
		model: 550,
		doors: 4,
		color: "Black",
	}

	fmt.Println(c1, c2)

}
