package chapter17

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindClosest(t *testing.T) {
	testData := []struct {
		name   string
		input  []string
		word1  string
		word2  string
		result *LocationPair
	}{
		{
			"1 distance",
			[]string{"how", "are", "we", "and", "you", "doing", "today", "and", "how", "many", "books", "are", "you", "reading", "in", "the", "morning"},
			"are",
			"you",
			&LocationPair{11, 12},
		},
		{
			"pair doesn't exist",
			[]string{"how", "are", "we", "and", "you", "doing", "today", "and", "how", "many", "books", "are", "you", "reading", "in", "the", "morning"},
			"are",
			"yousss",
			nil,
		},
		{
			"3 distance",
			[]string{"how", "are", "we", "and", "you", "doing", "today", "and", "how", "many", "books", "are", "you", "reading", "in", "the", "morning"},
			"how",
			"we",
			&LocationPair{0, 2},
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			dictionary := getWordLocations(data.input)
			assert.True(t, reflect.DeepEqual(data.result, FindClosest(data.word1, data.word2, dictionary)))
		})
	}
}
