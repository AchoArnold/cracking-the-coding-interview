package chapter10

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchStringExcludingEmpty(t *testing.T) {
	list := []string{
		"",
		"how",
		"fine",
		"",
		"",
		"car",
		"race",
		"",
	}

	t.Run("Testing when value is available", func(t *testing.T) {
		assert.Equal(t, 6, SearchStringExcludingEmpty(list, "race"))
	})

	t.Run("Testing when value is not available", func(t *testing.T) {
		assert.Equal(t, -1, SearchStringExcludingEmpty(list, "dsafdsds"))
	})
}
