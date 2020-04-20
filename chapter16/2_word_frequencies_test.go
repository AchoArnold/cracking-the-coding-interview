package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWordFrequency(t *testing.T) {
	testData := []struct {
		Name              string
		Input             string
		Word              string
		ExpectedFrequency int
	}{
		{
			Name:              "it returns zero when the word does not exist in the input",
			Input:             "how are you doing today and how is school",
			Word:              "man",
			ExpectedFrequency: 0,
		},
		{
			Name:              "it returns 1 when the word is at the beginning of the input",
			Input:             "are you doing today and how is school",
			Word:              "are",
			ExpectedFrequency: 1,
		},
		{
			Name:              "input is multi-case",
			Input:             "hoW are you doing today and HOW is school. What did you say yesterday how",
			Word:              "how",
			ExpectedFrequency: 3,
		},
		{
			Name:              "it returns 1 when the the word is at the end of the input",
			Input:             "how are you doing today and how is school man",
			Word:              "man",
			ExpectedFrequency: 1,
		},
	}

	for _, data := range testData {
		t.Run(data.Name, func(t *testing.T) {
			assert.Equal(t, data.ExpectedFrequency, GetWordFrequency(data.Input, data.Word))
		})
	}
}
