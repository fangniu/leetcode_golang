package main

//Given a binary tree, determine if it is height-balanced.
//
//For this problem, a height-balanced binary tree is defined as:
//
//a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
//Example 1:
//
//Given the following tree [3,9,20,null,null,15,7]:
//
//    3
//   / \
//  9  20
//     / \
//    15  7
//Return true.
//
//Example 2:
//
//Given the following tree [1,2,2,3,3,null,null,4,4]:
//
//      1
//     / \
//    2   2
//   / \
//  3   3
// / \
//4   4
//Return false.

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := depth(root.Left)
	right := depth(root.Right)
	if (left - right < 2) && (left - right) > -2 && isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}
	return false
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	a := depth(root.Left)
	b := depth(root.Right)
	if a < b {
		return b + 1
	} else {
		return a + 1
	}
}