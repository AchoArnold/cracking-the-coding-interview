package chapter4

func (root *BinaryTreeNode) CountPathsWithSum(targetSum int) int {
	return root.countPathsWithSumSub(targetSum, 0, map[int]int{})
}

func (root *BinaryTreeNode) countPathsWithSumSub(targetSum int, runningSum int, pathCount map[int]int) int {
	if root == nil {
		return 0
	}

	runningSum += root.Value.(int)

	sum := runningSum - targetSum
	totalPaths := 0
	if _, ok := pathCount[sum]; ok {
		totalPaths = pathCount[sum]
	}

	if runningSum == targetSum {
		totalPaths++
	}

	pathCount = incrementHashTable(pathCount, runningSum, 1)
	totalPaths += root.Left.countPathsWithSumSub(targetSum, runningSum, pathCount)
	totalPaths += root.Right.countPathsWithSumSub(targetSum, runningSum, pathCount)
	pathCount = incrementHashTable(pathCount, runningSum, -1)

	return totalPaths
}

func incrementHashTable(hashTable map[int]int, key int, delta int) map[int]int {
	newCount := delta
	if _, ok := hashTable[key]; ok {
		newCount += hashTable[key]
	}

	if newCount == 0 {
		delete(hashTable, key)
	} else {
		hashTable[key] = newCount
	}

	return hashTable
}
