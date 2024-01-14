package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		os.Exit(1)
	}

	mp := parsing(string(file))

	fmt.Printf("%v", mp)

	for k, v := range mp {
		fmt.Printf("%v : %v\n", string(k), v)
	}
}

func parsing(text string) map[rune]int {

	mp := make(map[rune]int)
	for _, ch := range text {
		mp[ch]++
	}
	return mp
}
