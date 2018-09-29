package main

//Given a binary search tree and the lowest and highest boundaries as L and R, trim the tree so that all its elements
// lies in [L, R] (R >= L). You might need to change the root of the tree, so the result should return the new root of
// the trimmed binary search tree.
//
//Example 1:
//Input:
//  1
// / \
//0   2
//
//L = 1
//R = 2
//
//Output:
// 1
//  \
//   2
//Example 2:
//Input:
//   3
//  / \
// 0   4
// \
//  2
//  /
// 1
//
//L = 1
//R = 3
//
//Output:
//     3
//    /
//   2
//  /
// 1

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 // 当 node.val > R, 说明修理后的树为node左边的子树
 // 当 node.val < L, 说明修理后的树为node右边的字树
 // ----------递归--------------------------
func trimBST(root *TreeNode, L int, R int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val > R {
		return trimBST(root.Left, L, R)
	}
	if root.Val < L {
		return trimBST(root.Right, L, R)
	}

	root.Left = trimBST(root.Left, L, R)
	root.Right = trimBST(root.Right, L, R)
	return root
}


