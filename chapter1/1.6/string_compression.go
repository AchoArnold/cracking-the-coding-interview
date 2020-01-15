package main

import (
	"strconv"
)

func stringCompression(inputString string) string {
	compressedString := ""
	currentString := string(inputString[0])
	counter := 0

	for index,character := range inputString   {
		if string(character) != currentString {
			compressedString += string(inputString[index - 1]) + strconv.Itoa(counter)
			currentString = string(inputString[index])
			counter = 1
 		} else {
 			counter++
		}
	}
	compressedString += string(inputString[len(inputString) -1]) + strconv.Itoa(counter)


	if len(compressedString) >= len(inputString) {
		return inputString
	}

	return compressedString
}