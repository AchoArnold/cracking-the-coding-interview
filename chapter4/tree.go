package chapter4

type BinaryTreeNode struct {
	Value interface{}
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

func DepthFirstSearchBinaryTree(node *BinaryTreeNode) []interface{} {
	if node == nil {
		return []interface{}{}
	}
	return append(DepthFirstSearchBinaryTree(node.Left), append([]interface{}{node.Value}, DepthFirstSearchBinaryTree(node.Right)...)...)
}
