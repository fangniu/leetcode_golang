package main

//Given an array where elements are sorted in ascending order, convert it to a height balanced BST.
//
//For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.
//
//
//Example:
//
//Given the sorted array: [-10,-3,0,5,9],
//
//One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:
//
// 	    0
//     / \
//   -3   9
//   /   /
// -10  5

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{nums[0], nil, nil}
	}
	if len(nums) == 2 {
		return &TreeNode{nums[1], &TreeNode{nums[0], nil, nil}, nil}
	}
	low := 0
	high := len(nums) - 1
	mid := low + (high-low)/2
	return &TreeNode{
		Val: nums[mid],
		Left: sortedArrayToBST(nums[low:mid]),
		Right: sortedArrayToBST(nums[mid+1:high+1]),
	}
}
