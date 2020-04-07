package chapter10

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	words := []string{
		"abut",
		"mace",
		"race",
		"tuba",
		"care",
	}

	GroupAnagrams(words)

	assert.True(t, reflect.DeepEqual([]string{"abut", "tuba", "mace", "race", "care"}, words))
}
