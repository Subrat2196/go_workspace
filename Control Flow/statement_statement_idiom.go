/* Statement; statement idiom  -> it means we can have a statement and
ok now check statement idiom inside if */

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := 45
	if z := 2 * rand.Intn(45); z >= x {
		fmt.Printf("The value of z is %v and x is %v and z is greater or equal to x \n", z, x)
	} else {
		fmt.Printf("The value of z is %v and x is %v and z is smaller than x \n", z, x)
	}
}
