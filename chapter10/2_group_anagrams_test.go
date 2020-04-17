package chapter10

import (
	"reflect"
	"testing"
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

	if !reflect.DeepEqual([]string{"abut", "tuba", "mace", "race", "care"}, words) && !reflect.DeepEqual([]string{"race", "care", "abut", "tuba", "mace"}, words) {
		t.Errorf("%#+v  is not grouped", words)
	}
}
