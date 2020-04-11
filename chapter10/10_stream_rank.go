package chapter10

// RankNode is a data structure for left size aware binary search tree
type RankNode struct {
	leftSize int
	left     *RankNode
	right    *RankNode
	data     int
}

// NewRankNode creates a new tree node
func NewRankNode(data int) *RankNode {
	return &RankNode{
		leftSize: 0,
		left:     nil,
		right:    nil,
		data:     data,
	}
}

func (node *RankNode) insert(d int) {
	if d <= node.data {
		if node.left != nil {
			node.left.insert(d)
		} else {
			node.left = NewRankNode(d)
			node.leftSize++
		}
	} else {
		if node.right != nil {
			node.right.insert(d)
		} else {
			node.right = NewRankNode(d)
		}
	}
}

func (node *RankNode) getRank(d int) int {
	if d == node.data {
		return node.leftSize
	}

	if d < node.data {
		if node.left == nil {
			return -1
		}
		return node.left.getRank(d)
	}

	rightRank := -1
	if node.right != nil {
		rightRank = node.right.getRank(d)
	}

	if rightRank == -1 {
		return -1
	}

	return node.leftSize + 1 + rightRank
}
