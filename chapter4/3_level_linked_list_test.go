package chapter4

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateLevelLinkedList(t *testing.T) {
	expected := []int{1, 2, 3, 4, 5}

	tree := GetBinaryTreeFromSortedArray(expected)

	listsArray := CreateLevelLinkedList(tree)

	assert.Equal(t, 3, len(listsArray))

	var list1Items []int
	for item := listsArray[0].Front(); item != nil; item = item.Next() {
		list1Items = append(list1Items, item.Value.(*BinaryTreeNode).Value.(int))
	}
	assert.Equal(t, []int{3}, list1Items)

	var list2Items []int
	for item := listsArray[1].Front(); item != nil; item = item.Next() {
		list2Items = append(list2Items, item.Value.(*BinaryTreeNode).Value.(int))
	}
	assert.Equal(t, []int{2, 5}, list2Items)

	list2Items = []int{}
	for item := listsArray[2].Front(); item != nil; item = item.Next() {
		list2Items = append(list2Items, item.Value.(*BinaryTreeNode).Value.(int))
	}
	assert.Equal(t, []int{1, 4}, list2Items)
}
