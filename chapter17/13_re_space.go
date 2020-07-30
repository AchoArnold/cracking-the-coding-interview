package chapter17

import (
	"math"
)

// ParseResult stores the result of parsing a string
type ParseResult struct {
	invalid int
	parsed  string
}

func newParseResult(inv int, p string) *ParseResult {
	return &ParseResult{
		invalid: inv,
		parsed:  p,
	}
}

// BestSplit splits a string on dictionary words and returns the best parse.
func BestSplit(dictionary map[string]struct{}, sentence string) string {
	cache := make([]*ParseResult, len(sentence))
	result := splitString(dictionary, sentence, 0, cache)

	if result == nil {
		return ""
	}

	return result.parsed
}

func splitString(dictionary map[string]struct{}, sentence string, start int, cache []*ParseResult) *ParseResult {
	if start >= len(sentence) {
		return newParseResult(0, "")
	}

	if cache[start] != nil {
		return cache[start]
	}

	bestInvalid := math.MaxInt32
	bestParsing := ""
	partial := ""
	index := start

	for index < len(sentence) {
		char := sentence[index]
		partial += string(char)

		invalid := len(partial)
		if _, ok := dictionary[partial]; ok {
			invalid = 0
		}

		if invalid < bestInvalid { // short circuit
			// Recurse, putting a space after this character.
			// If this is a better than the current best option, replace the best option.
			result := splitString(dictionary, sentence, index+1, cache)
			if (invalid + result.invalid) < bestInvalid {
				bestInvalid = invalid + result.invalid

				bestParsing = partial + " " + result.parsed

				if bestInvalid == 0 { // short circuit
					break
				}
			}
		}

		index++
	}

	cache[start] = newParseResult(bestInvalid, bestParsing)
	return cache[start]
}
