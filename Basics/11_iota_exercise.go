package main

import "fmt"

type ByteSize int

func main() {
	const (
		_           = iota
		KB ByteSize = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
	)

	fmt.Println(KB, MB, GB, TB, PB, EB)

}
