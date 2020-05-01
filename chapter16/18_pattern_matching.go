package chapter16

import (
	"strings"
)

// CheckPatternMatchesBruteForce checks if value matches pattern
func CheckPatternMatchesBruteForce(pattern, value string) bool {
	if len(pattern) == 0 {
		return len(value) == 0
	}

	size := len(value)
	for mainSize := 0; mainSize < size; mainSize++ {
		main := value[0:mainSize]
		for altStart := mainSize; altStart <= size; altStart++ {
			for altEnd := altStart; altEnd < size; altEnd++ {
				alt := value[altStart:altEnd]
				candidate := buildFromPatten(pattern, main, alt)
				if candidate == value {
					return true
				}
			}
		}
	}

	return false
}

// CheckPatternMatchesOptimized checks if value matches pattern
func CheckPatternMatchesOptimized(pattern, value string) bool {
	if len(pattern) == 0 {
		return len(value) == 0
	}

	mainChar := int32(pattern[0])
	altChar := 'a'
	if mainChar == 'a' {
		altChar = 'b'
	}

	size := len(value)

	countOfMain := countOf(pattern, mainChar)
	countOfAlt := len(pattern) - countOfMain
	firstAlt := strings.Index(pattern, string(altChar))
	maxMainSize := size / countOfMain

	for mainSize := 0; mainSize <= maxMainSize; mainSize++ {
		remainingLength := size - mainSize*countOfMain
		if countOfAlt == 0 || remainingLength%countOfAlt == 0 {
			altIndex := firstAlt * mainSize
			altSize := 0
			if countOfAlt != 0 {
				altSize = remainingLength / countOfAlt
			}

			if matches(pattern, value, mainSize, altSize, altIndex) {
				return true
			}
		}
	}

	return false
}

func matches(pattern string, value string, mainSize int, altSize int, firstAlt int) bool {
	stringIndex := mainSize
	for i := 1; i < len(pattern); i++ {
		size := altSize
		if pattern[0] == pattern[i] {
			size = mainSize
		}

		offset := firstAlt
		if pattern[i] == pattern[0] {
			offset = 0
		}

		if !isEqual(value, offset, stringIndex, size) {
			return false
		}

		stringIndex += size
	}

	return true
}

func isEqual(s1 string, offset1 int, offset2 int, size int) bool {
	for i := 0; i < size; i++ {
		if s1[offset1+i] != s1[offset2+i] {
			return false
		}
	}

	return true
}
func countOf(pattern string, char int32) int {
	count := 0
	for i := 0; i < len(pattern); i++ {
		if int32(pattern[i]) == char {
			count++
		}
	}

	return count
}

func buildFromPatten(pattern, main, alt string) string {
	sb := new(strings.Builder)
	first := int32(pattern[0])

	for _, char := range pattern {
		if char == first {
			sb.WriteString(main)
		} else {
			sb.WriteString(alt)
		}
	}

	return sb.String()
}
