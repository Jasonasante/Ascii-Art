package main

import (
	"fmt"
	"os"
)

func main() {
	words := os.Args[1]
	runes := []rune(words)
	fmt.Println(string(runes))
}
