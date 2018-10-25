package main
//
//Given a binary tree, find the length of the longest path where each node in the path has the same value. This path
// may or may not pass through the root.
//
//Note: The length of path between two nodes is represented by the number of edges between them.
//
//Example 1:
//
//Input:
//
//    5
//   / \
//  4   5
// / \   \
//1   1   5
//Output:
//
//2
//Example 2:
//
//Input:
//
//    1
//   / \
//  4   5
// / \   \
//4   4   5
//Output:
//
//2
//Note: The given binary tree has not more than 10000 nodes. The height of the tree is not more than 1000.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestUnivaluePath(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sub := max(longestUnivaluePath(root.Left), longestUnivaluePath(root.Right))
	return max(sub, longestUnivaluePathHelper(root.Left, root.Val) + longestUnivaluePathHelper(root.Right, root.Val))
}

func longestUnivaluePathHelper(root *TreeNode, parent int) int {
	if root == nil || root.Val != parent {
		return 0
	}
	return 1 + max(longestUnivaluePathHelper(root.Left, root.Val), longestUnivaluePathHelper(root.Right, root.Val))
}


func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}