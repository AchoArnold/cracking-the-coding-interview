package chapter16

import "unicode"

/* Mapping of digits to letters. */
var t9Letters = [][]int32{
	nil,
	nil,
	{'a', 'b', 'c'},
	{'d', 'e', 'f'},
	{'g', 'h', 'i'},
	{'j', 'k', 'l'},
	{'m', 'n', 'o'},
	{'p', 'q', 'r', 's'},
	{'t', 'u', 'v'},
	{'w', 'x', 'y', 'z'},
}

// GetValidT9Words returns words in dictionary having a t9 encoding
func GetValidT9Words(number string, dictionary map[string]string) []string {
	return getValidWords(number, 0, "", dictionary, []string{})
}

func getValidWords(number string, index int, prefix string, wordSet map[string]string, results []string) []string {
	// if it's a compute word, print it.
	_, ok := wordSet[prefix]
	if index == len(number) && ok {
		return append(results, prefix)
	}

	// get characters that match this digit
	var letters []int32
	if index < len(number) {
		digit := int32(number[index])
		letters = getT9Chars(digit)
	}

	// go through all remaining options
	if letters != nil {
		for _, letter := range letters {
			results = getValidWords(number, index+1, prefix+string(letter), wordSet, results)
		}
	}

	return results
}

func getT9Chars(digit int32) []int32 {
	if !unicode.IsDigit(digit) {
		return nil
	}
	return t9Letters[digit-'0']
}
