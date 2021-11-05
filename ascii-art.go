package main

import (
	"io/ioutil"
	"os"
)

func main() {
	words := os.Args[1]
	runes := []rune(words)
	emptySLice := [8][]byte{}
	for _, char := range runes {
		content, _ := ioutil.ReadFile(os.Args[2])
	}
	// num := 9
	// matrix := make([][]byte, num)
	// for i := 0; i < num; i++ {
	// 	matrix[i] = make([]byte, num)
	// 	vector := make([]byte, num)
	// 	for j := 0; j < num; j++ {
	// 		vector[j] = i*num + j
	// 		matrix[i] = append(matrix[i], vector[j])
	// 	}
	// }
}
