package main

import "fmt"

func main() {
	mp := make(map[string][]string)

	mp["bond_james"] = []string{"shaken", "not stirred", "martinis", "fast cars"}
	mp["moneypenny_jenny"] = []string{"intelligence", "literature", "computer science"}
	mp["no_dr"] = []string{"cats", "ice cream", "sunsets"}
	mp["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	delete(mp, "bond_james")

	for key, val := range mp {
		fmt.Println(key, "  ", val)
		for i, v := range val {
			fmt.Print(i, "  ", v, "\n")
		}
	}

}
