package chapter4

import "container/list"

func (root *BinaryTreeNode) AllSequences() *list.List {
	result := list.New()

	if root == nil {
		result.PushBack(list.New())
		return result
	}

	prefix := list.New()
	prefix.PushBack(root.Value)

	//Recurse on left and right subtrees
	leftSeq := root.Left.AllSequences()
	rightSeq := root.Right.AllSequences()

	for left := leftSeq.Front(); left != nil; left = left.Next() {
		for right := rightSeq.Front(); right != nil; right = right.Next() {
			weaved := list.New()
			weaveLists(left.Value.(*list.List), right.Value.(*list.List), weaved, prefix)

			result.PushBackList(weaved)

		}
	}

	return result
}

func weaveLists(first *list.List, second *list.List, results *list.List, prefix *list.List) {
	// One list is empty. Add reminder to a clone prefix and store result.
	if first.Len() == 0 || second.Len() == 0 {
		result := list.New()
		for i := prefix.Front(); i != nil; i = i.Next() {
			result.PushBack(i.Value)
		}

		result.PushBackList(first)
		result.PushBackList(second)
		results.PushBack(result)

		return
	}

	// Recurse with head first added to the prefix. Removing the head will damage
	// first so we'll nee to put it back where we found it afterwards
	headFirst := first.Front()
	first.Remove(headFirst)
	prefix.PushBack(headFirst.Value)
	weaveLists(first, second, results, prefix)
	lastItem := prefix.Back()
	prefix.Remove(lastItem)
	first.PushFront(headFirst.Value)

	// Do the same thing with second, damaging and then restoring the list
	headSecond := second.Front()
	second.Remove(headSecond)
	prefix.PushBack(headSecond.Value)
	weaveLists(first, second, results, prefix)
	lastPrefix := prefix.Back()
	prefix.Remove(lastPrefix)
	second.PushFront(headSecond.Value)

}
