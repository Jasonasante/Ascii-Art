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
		// this takes the argument that we are printing and seperates them into a []string via \n
		// this will then therefore automatically will print each []string on a new line.
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
	// into a slice of string
	for scanner.Scan() {
		// i think scanner.Scan
		if scanner.Text() == "" && count != 0 {
			// if inside that slice is empty, and count != 0 then append tempAscii (which at this point will be filled the buuble text)
			// as we are dealing with else if statement, this has to come first
			// this also means that the very first line of the file is skipped since at that point, count will == 0
			asciiSlice = append(asciiSlice, tempAscii)
			tempAscii = nil
			// then make tempAscii empty, ready to be filled by the next set of bubble text
		} else if count != 0 {
			tempAscii = append(tempAscii, scanner.Text())
			// this appends what is inside scanner.Text since it contains texts
		}
		count = 1
	}
	asciiSlice = append(asciiSlice, tempAscii)
	// at the end of the file, there is no new line and therefore the last bubble text will be appended to tempAscii, however, that will not
	// be appended to asciiSLice due to the condition ' if scanner.Text() == "" && count != 0'
	// so we have to append the remaining tempAscii outside of the loop.
	tempOutput = buildOutputArr()
	// this makes tempOutput 8 empty slices of string
	for _, str := range inputString {
		for _, srune := range str {
			for i, asciiRune := range asciiString {
				// remember asciiString is in the same order as asciiSlice (which is the bubble text)
				if srune == asciiRune {
					// so when the srune==asciiRune, then take the index of asciiRune, the bubble text equivalent to tempOutput
					// which is the empty 8 []string
					tempOutput = addAscii(asciiSlice[i], tempOutput)
				}
			}
		}
		outputAscii = append(outputAscii, tempOutput)
		// then append that bubble text to output ascii,
		tempOutput = buildOutputArr()
		// then make tempOutput back into 8 empty []string
	}
	for _, outputArr := range outputAscii {
		for _, outputStr := range outputArr {
			fmt.Println(outputStr)
		}
	}
}
