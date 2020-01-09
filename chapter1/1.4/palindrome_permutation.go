package main

import (
	"strings"
)

func isPalindromePermutation(inputString string) bool {
	spaceCharacter := " "

	inputString = strings.ToLower(inputString)

	stringRune := map[rune]int{}

	for _, character := range inputString  {
		if string(character) != spaceCharacter {
			stringRune[character]++
		}
	}

	hasOddNumber := false

	for _, count := range stringRune  {
		if count %2 != 0 {
			if hasOddNumber == true {
				return false
			}

			hasOddNumber = true
		}
	}

	return true
}

