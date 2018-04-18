package main


//Given a binary tree, find its maximum depth.
//
//The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
//For example:
//Given binary tree [3,9,20,null,null,15,7],
//
//	  3
//	 / \
//	9   20
//	   /  \
//    15   7
//return its depth = 3.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	if root.Left == nil {
		return maxDepth(root.Right) + 1
	}
	if root.Right == nil {
		return maxDepth(root.Left) + 1
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left < right {
		return right + 1
	} else {
		return left + 1
	}
}

func main() {
	maxDepth(nil)
}