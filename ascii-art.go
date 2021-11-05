package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func buildOutputArr() []string {
	var strarr []string
	for i := 0; i < 8; i++ {
		strarr = append(strarr, "")
	}
	return strarr
}

func addAscii(asciiChar []string, outputArr []string) []string {
	for i, str := range asciiChar {
		outputArr[i] += str
	}
	return outputArr
}

func main() {
	var asciiSlice [][]string
	var tempAscii []string
	var outputAscii [][]string
	var tempOutput []string
	var inputString []string
	count := 0
	asciiString := ""
	if len(os.Args) == 2 {
		inputString = strings.Split(os.Args[1], "\\n")
	} else {
		fmt.Println("to few or to many arguments")
		os.Exit(0)
	}
	for i := 32; i <= 126; i++ {
		asciiString += string(rune(i))
		// this makes the index of asciiString in the same order as that of the banner template files.
		// so asciiString[0]== " " and asciiString[1]== "!" and so on
	}
	file, _ := os.Open("standard.txt")
	scanner := bufio.NewScanner(file)
	// bufio.NewScanner separates the file (in this case standard.txt) line by line
	fmt.Println(scanner)
	for scanner.Scan() {
		if scanner.Text() == "" && count != 0 {
			asciiSlice = append(asciiSlice, tempAscii)
			tempAscii = nil
		} else if count != 0 {
			tempAscii = append(tempAscii, scanner.Text())
		}
		count = 1
	}
	asciiSlice = append(asciiSlice, tempAscii)
	tempOutput = buildOutputArr()
	for _, str := range inputString {
		for _, srune := range str {
			for i, asciiRune := range asciiString {
				if srune == asciiRune {
					tempOutput = addAscii(asciiSlice[i], tempOutput)
				}
			}
		}
		outputAscii = append(outputAscii, tempOutput)
		tempOutput = buildOutputArr()
	}
	for _, outputArr := range outputAscii {
		for _, outputStr := range outputArr {
			fmt.Println(outputStr)
		}
	}
}
