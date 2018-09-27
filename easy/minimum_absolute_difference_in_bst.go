package main

import "math"

//Given a binary search tree with non-negative values, find the minimum absolute difference between values of any two nodes.
//
//Example:
//
//Input:
//
//1
// \
//  3
// /
//2
//
//Output:
//1
//
//Explanation:
//The minimum absolute difference is 1, which is the difference between 2 and 1 (or between 2 and 3).
//Note: There are at least two nodes in this BST.
//


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func getMinimumDifference(root *TreeNode) int {
	var (
		minDiff = math.MaxInt32
		prev    = -1
	)
	traverseInOrder(root, &minDiff, &prev)
	return minDiff
}
func traverseInOrder(root *TreeNode, minDiff *int, prev *int) {
	if root == nil {
		return
	}
	traverseInOrder(root.Left, minDiff, prev)

	if *prev != -1 {
		*minDiff = minValue(*minDiff, root.Val - *prev)
	}
	*prev = root.Val
	traverseInOrder(root.Right, minDiff, prev)
}

func minValue(a int, b int)  int {
	if a > b {
		return b
	}
	return a
}