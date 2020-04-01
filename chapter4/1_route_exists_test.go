package chapter4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraph_RouteExists(t *testing.T) {
	graph := Graph{
		[]GraphNode{
			{
				Value:    "0",
				Children: StringToInterfaceSlice([]string{"1", "4", "5"}),
			},
			{
				Value:    "1",
				Children: StringToInterfaceSlice([]string{"4", "3"}),
			},
			{
				Value:    "2",
				Children: StringToInterfaceSlice([]string{"1"}),
			},
			{
				Value:    "3",
				Children: StringToInterfaceSlice([]string{"2", "4"}),
			},
			{
				Value:    "4",
				Children: []interface{}{},
			},
		},
	}

	t.Run("It returns true when a route exists", func(t *testing.T) {
		assert.True(t, graph.RouteExists("0", "2"))
		assert.True(t, graph.RouteExists("1", "1"))
	})

	t.Run("It returns false when a route does not exists", func(t *testing.T) {
		assert.False(t, graph.RouteExists("1", "0"))
		assert.False(t, graph.RouteExists("3", "0"))
	})
}
