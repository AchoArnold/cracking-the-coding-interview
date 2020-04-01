package chapter4

import (
	"container/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTreeNode_AllSequences(t *testing.T) {
	tree := GetBinaryTreeFromSortedArray([]int{1, 2, 3})

	seq := tree.AllSequences()

	var res [][]interface{}
	for i := seq.Front(); i != nil; i = i.Next() {
		var inner []interface{}
		for j := i.Value.(*list.List).Front(); j != nil; j = j.Next() {
			inner = append(inner, j.Value)
		}

		res = append(res, inner)
	}

	assert.Equal(t, [][]interface{}{[]interface{}{2, 1, 3}, []interface{}{2, 3, 1}}, res)
}
