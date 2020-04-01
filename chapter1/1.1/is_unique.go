package chapter1

func isUnique(inputString string) bool {
	for index, character := range inputString {
		subStringStartIndex := index + 1
		for _, subCharacter := range inputString[subStringStartIndex:] {
			if subCharacter == character {
				return false
			}

		}
	}

	return true
}

func isUniqueMap(inputString string) bool {
	runeSet := map[rune]struct{}{}

	for _, character := range inputString {
		if _, characterExists := runeSet[character]; !characterExists {
			runeSet[character] = struct{}{}
		} else {
			return false
		}

	}

	return true
}
