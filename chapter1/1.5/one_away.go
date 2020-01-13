package main

func oneAway(inputString1 string, inputString2 string) bool {
	stringAsRune := map[rune]int{}

	for _,character := range inputString1 {
		stringAsRune[character]++
	}

	invalidFound := false

	for _,character := range inputString2 {
		if _, exists := stringAsRune[character]; !exists {
			if invalidFound {
				return false
			} else  {
				invalidFound = true
			}
		} else {
			stringAsRune[character]--
		}
	}

	invalidFound = false
	for _, value := range stringAsRune {
		if value != 0 {
			if value < -1 || value > 1 || invalidFound {
				return  false
			} else {
				invalidFound = true
			}
		}
	}

	return true
}