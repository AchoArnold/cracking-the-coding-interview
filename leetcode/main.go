package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	println(simplifyPath("/a/./b/../../c/"))
}

func simplifyPath(path string) string {
	paths := []string{}

	start := 0

	for start < len(path) {
		nextPath := getNextPath(path, start)

		if nextPath == ".." {
			if len(paths) > 0 {
				paths = paths[0: len(paths) -1]
			}
		} else if nextPath != "." && nextPath != "" {
			paths = append(paths, nextPath)
		}

		start += len(nextPath) + 1
	}

	if len(paths) == 0 {
		return "/"
	}

	result := ""
	for _, val  := range paths {
		result += "/" + val
	}

	return result
}

func getNextPath(path string, start int) (res string){
	temp := path[start:]
	result := ""
	for _, char := range temp {
		if char != '/' {
			result += string(char)
		} else {
			return result
		}
	}
	return result
}

// Function to insert nodes in level order
func InsertLevelOrder(array []interface{}, i int) *TreeNode {
	if i >= len(array) || array[i] == nil {
		return nil
	}

	node := &TreeNode{
		Val:   array[i].(int),
		Left:  nil,
		Right: nil,
	}

	node.Left = InsertLevelOrder(array, 2 *i + 1)
	node.Right = InsertLevelOrder(array, 2* i + 2)

	return node
}

func (root *TreeNode) Print() {
	if root != nil {
		root.Left.Print()
		fmt.Printf("%d ", root.Val)
		root.Right.Print()
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isCousins(root *TreeNode, x int, y int) bool {
	if root == nil {
		return false
	}

	nodes := []*TreeNode{root}

	for len(nodes) != 0 {
		var	newQueue []*TreeNode
		xFound := false
		yFound := false

		for _, node := range nodes {
			localX := false
			localY := false
			if nodeHasVal(node, x) {
				localX = true
			}

			if nodeHasVal(node, y) {
				localY = true
			}

			if localX && localY {
				return false
			}

			if localX {
				xFound = true
			}

			if localY {
				yFound = true
			}

			if xFound && yFound {
				return true
			}

			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}

			if root.Right != nil {
				newQueue = append(newQueue, node.Right)

			}
		}

		if xFound != yFound {
			return false
		}

		nodes = newQueue
	}

	return false
}

func nodeHasVal(node *TreeNode, val int) bool {
	if node == nil {
		return false
	}

	if node.Left != nil && node.Left.Val == val {
		return true
	}

	if node.Right != nil && node.Right.Val == val {
		return true
	}

	return false
}
