package chapter8

func GetPermutations(str string) (result []string) {
	if len(str) == 0 {
		return []string{""}
	}

	first := str[0]
	charArray := []string{string(first)}

	remainder := ""
	if len(str) > 1 {
		remainder = str[1:]
	}

	for _, char := range remainder {
		charArray = insertChar(char, charArray)
	}

	return charArray
}

func insertChar(char int32, charArray []string) []string {
	result := []string{}
	for _, word := range charArray {
		for index, _ := range word {
			result = append(result, insertCharAt(word, char, index))
		}
		result = append(result, word+string(char))
	}

	return result
}

func insertCharAt(word string, char int32, index int) string {
	start := word[0:index]
	end := word[index:]

	return start + string(char) + end
}
