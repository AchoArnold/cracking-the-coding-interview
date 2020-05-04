package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testDataPatternCheck = []struct {
	name, pattern, value string
	expected             bool
}{
	{"empty string", "", "", true},
	{"matching string", "aabab", "catcatgocatgo", true},
	{"non-matching string", "aabab", "bcatbcatgocatgo", false},
}

func TestCheckPatternMatchesBruteForce(t *testing.T) {
	for _, data := range testDataPatternCheck {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expected, CheckPatternMatchesBruteForce(data.pattern, data.value))
		})
	}
}

func TestCheckPatternMatchesOptimized(t *testing.T) {
	for _, data := range testDataPatternCheck {
		t.Run(data.name, func(t *testing.T) {
			assert.Equal(t, data.expected, CheckPatternMatchesOptimized(data.pattern, data.value))
		})
	}
}
