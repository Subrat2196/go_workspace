// what does a variadic parameter means
// in a function what can differ is the number of arguments that you pass
// we already know a function that can take unlimited number of arguments i.e the Println function
// func Println(a,.....any)(n int,err error) -> here any means empty interface
// therefore Println function is a type of variadice parameter
// when we pass an unlimited number of parameters let's say e.g integers, so wo will change them to
// slice of int

// the final parameter should be like ii ...int , we can have other fixed parameters before it also
// this final parameter is called variadic parameter
// a variadic parameter will have type prefixed with ...

package main

import "fmt"

func main() {

	a := sum(5, "1", "2", "3", "4", "5")
	fmt.Println(a)

}

func sum(n int, ss ...string) string {

	fmt.Println(ss)
	fmt.Printf("%T\n", ss)

	for _, val := range ss {
		fmt.Sprintln(val)
	}

	return "Hello"
}
