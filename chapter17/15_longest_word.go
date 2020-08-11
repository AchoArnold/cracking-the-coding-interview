package chapter17

import (
	"sort"
)

// GetLongestWord returns the longest word in a list made up of sub words.
func GetLongestWord(words []string) string {
	cache := make(map[string]bool, len(words))
	for _, word := range words {
		cache[word] = true
	}

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	for _, word := range words {
		if canBuildWord(word, true, cache) {
			return word
		}
	}

	return ""
}

func canBuildWord(word string, isOriginalWord bool, cache map[string]bool) bool {
	if _, ok := cache[word]; ok && !isOriginalWord {
		return cache[word]
	}

	for i := 1; i < len(word); i++ {
		left := word[0:i]
		right := word[i:]

		if val, ok := cache[left]; ok && val && canBuildWord(right, false, cache) {
			return true
		}
	}

	cache[word] = false
	return false
}
