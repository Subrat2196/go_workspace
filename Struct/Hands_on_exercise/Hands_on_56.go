package main

import "fmt"

type anony struct {
	first      string
	friends    map[string]int
	fav_drinks []string
}

func main() {

	a1 := struct {
		first      string
		friends    map[string]int
		fav_drinks []string
	}{
		first: "Subrat",
		friends: map[string]int{
			"Rahul":  1,
			"shivam": 2,
		},
		fav_drinks: []string{
			"martini", "cosmo",
		},
	}

	fmt.Println(a1)

}
