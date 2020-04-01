package chapter4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindBuildOrder(t *testing.T) {
	t.Run("Test it returns the correct build order when there is a valid build order", func(t *testing.T) {
		projects := FindBuildOrder(
			[]string{"a", "b", "c", "d", "e", "f"},
			[][]string{
				{"a", "d"},
				{"f", "b"},
				{"b", "d"},
				{"f", "a"},
				{"d", "c"},
			},
		)

		expectedOrder := []string{"e", "f", "b", "a", "d", "c"}

		var actualOrder []string
		for _, project := range projects {
			actualOrder = append(actualOrder, project.GetName())
		}

		assert.Equal(t, expectedOrder, actualOrder)

	})

	t.Run("Test it throws a panic when there's a circular dependency", func(t *testing.T) {
		assert.Panics(t, func() {
			FindBuildOrder(
				[]string{"a", "b", "c", "d", "e", "f"},
				[][]string{
					{"a", "d"},
					{"f", "b"},
					{"b", "d"},
					{"d", "b"},
					{"d", "c"},
				},
			)
		})

	})
}
