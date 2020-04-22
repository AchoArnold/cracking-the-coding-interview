package chapter10

import (
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	words := []string{
		"abut",
		"race",
		"tuba",
		"care",
	}

	GroupAnagrams(words)

	if !reflect.DeepEqual([]string{"abut", "tuba", "race", "care"}, words) && !reflect.DeepEqual([]string{"race", "care", "abut", "tuba"}, words) {
		t.Errorf("%#+v  is not grouped", words)
	}
}
