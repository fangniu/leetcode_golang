package main

import "fmt"

//Given a binary tree, return the tilt of the whole tree.
//
//The tilt of a tree node is defined as the absolute difference between the sum of all left subtree node values and the
// sum of all right subtree node values. Null node has tilt 0.
//
//The tilt of the whole tree is defined as the sum of all nodes' tilt.
//
//Example:
//Input:
//   1
// /   \
//2     3
//Output: 1
//Explanation:
//Tilt of node 2 : 0
//Tilt of node 3 : 0
//Tilt of node 1 : |2-3| = 1
//Tilt of binary tree : 0 + 0 + 1 = 1
//Note:
//
//The sum of node values in any subtree won't exceed the range of 32-bit integer.
//All the tilt values won't exceed the range of 32-bit integer.


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	tilt := 0
	traverse(root, &tilt)
	return tilt
}

func traverse(root *TreeNode, tilt *int) int {
	if root == nil {
		return 0
	}
	left := traverse(root.Left, tilt)
	right := traverse(root.Right, tilt)
	*tilt = *tilt + diff(left, right)
	return left + right + root.Val
}

func diff(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	fmt.Println(findTilt(nil))
}