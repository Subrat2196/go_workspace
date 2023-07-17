// Anonymous structs are used when we dont want to create an entry of a
//particulartype in the code, we can use anonymous struct

package main

import "fmt"

type person struct {
	first  string
	second string
	age    int
}

type secret_service struct {
	person
	first string
	ltk   bool
	guns  int
}

func main() {

	p1 := struct {
		first  string
		second string
		age    int
	}{
		first:  "Subrat",
		second: "Bahuguna",
		age:    27,
	}

	// sa1 := secret_service{
	// 	person: person{
	// 		first:  "saint",
	// 		second: "bahuguna",
	// 		age:    27,
	// 	},
	// 	ltk:  true,
	// 	guns: 3,
	// }

	/* If we replace person above with the definition of person in original struct then we have
	anonymous struct i.e
	 person: (replace this) {
		first:"saint"

	}
	*/

	// sa1 := secret_service{
	// 	person: person{
	// 		first:  "saint",
	// 		second: "bahuguna",
	// 		age:    27,
	// 	},
	// 	ltk:  true,
	// 	guns: 3,
	// }

	fmt.Printf("The person struct is %v \n and type of sa1 is %#T \n", p1, p1)

}
