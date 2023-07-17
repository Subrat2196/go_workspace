// Word Frequency
package main

import "fmt"

func main() {

	words := []string{"hello", "there", "i", "am", "subrat", "subrat", "subrat", "subrat", "subrat", "subrat", "subrat", "there", "there", "there", "there", "there", "there", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "hello", "i", "i", "i", "i", "i", "i", "am", "am", "am", "am", "am"}

	mp_freq := make(map[string]int)

	fmt.Println(words)

	for _, v := range words {
		mp_freq[v]++
	}

	for key, val := range mp_freq {
		fmt.Println(key, "  ", val)
	}
}
