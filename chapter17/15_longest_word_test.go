package chapter17

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLongestWord(t *testing.T) {
	testData := []struct {
		name   string
		input  []string
		output string
	}{
		{
			"Empty string",
			[]string{},
			"",
		},
		{
			"two longest words",
			[]string{"how", "are", "you", "doing", "moon", "light", "moonlight", "moonlighting"},
			"moonlight",
		},
		{
			"3 longest word",
			[]string{"how", "are", "you", "doing", "today", "howdoingyou", "fine"},
			"howdoingyou",
		},
	}

	for _, data := range testData {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.output, GetLongestWord(data.input))
		})
	}
}
