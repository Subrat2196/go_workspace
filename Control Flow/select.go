// Select statements are used when we want to write concurrent codes and dealing with channels
// Select statements are visually similar to switch statements
// Select statament chooses which set of possible send or receive operations will proceed
// in select all the cases refer to communication operations
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Creating a channel that can transfer integers between goroutines
	ch1 := make(chan int)
	ch2 := make(chan int)

	// collecting a random time duration between [0,250)
	d1 := time.Duration(rand.Int63n(250))
	d2 := time.Duration(rand.Int63n(250))

	go func() {
		time.Sleep(d1 * time.Millisecond) // making that time duration to milliseconds
		ch1 <- 41                         // this go routine will push a value 41 to ch1 whenever it finishes sleeping
	}()

	go func() {
		time.Sleep(d2 * time.Millisecond)
		ch2 <- 42
	}()

	select {
	case v1 := <-ch1:
		fmt.Println("the value from channel 1", v1)

	case v2 := <-ch2:
		fmt.Println("value from channel 2", v2)
	}

}
