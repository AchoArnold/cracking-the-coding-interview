package chapter10

import (
	"sort"
	"strings"
)

// GroupAnagrams groups anagrams in an array in place.
func GroupAnagrams(array []string) {
	cache := map[string][]string{}

	// group words by anagram
	for _, word := range array {
		key := strings.Split(word, "")
		sort.Strings(key)
		keys := strings.Join(key, "")
		cache[keys] = append(cache[keys], word)
	}

	index := 0
	for _, valArray := range cache {
		for _, val := range valArray {
			array[index] = val
			index++
		}
	}
}
