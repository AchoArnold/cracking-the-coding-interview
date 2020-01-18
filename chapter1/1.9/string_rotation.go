package main

import "strings"

func isStringRotation(inputString1 string, inputString2 string) bool {
	if len(inputString2) != len(inputString1) {
		return false
	}

	inputString2 = inputString2 + inputString2

	return strings.Contains(inputString2, inputString1)
}
