package chapter4

func (root *BinaryTreeNode) ContainsTree(childRoot *BinaryTreeNode) bool {
	if childRoot == nil {
		return true
	}
	return root.HasSubTree(childRoot)
}

func (root *BinaryTreeNode) HasSubTree(childRoot *BinaryTreeNode) bool {
	if root == nil {
		return false
	}

	if root.Value == childRoot.Value && root.MatchTree(childRoot) {
		return true
	}

	return root.Left.HasSubTree(childRoot) || root.Right.HasSubTree(childRoot)
}

func (root *BinaryTreeNode) MatchTree(childRoot *BinaryTreeNode) bool {
	if root == nil && childRoot == nil {
		return true
	}

	if root == nil || childRoot == nil {
		return false
	}

	if root.Value != childRoot.Value {
		return false
	}

	return root.Left.MatchTree(childRoot.Left) && root.Right.MatchTree(childRoot.Right)
}
