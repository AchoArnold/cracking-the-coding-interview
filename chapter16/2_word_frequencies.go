package chapter16

import "strings"

// GetWordFrequency returns how many time a string exists in an input
func GetWordFrequency(input string, word string) int {
	input = strings.ToLower(input)
	word = strings.Trim(strings.ToLower(word), " ")
	words := strings.Split(input, " ")

	frequency := 0
	for _, inputWord := range words {
		if inputWord == word {
			frequency++
		}
	}

	return frequency
}
