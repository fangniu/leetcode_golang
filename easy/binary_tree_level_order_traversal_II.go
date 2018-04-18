package main

import (
	"fmt"
)

//Given a binary tree, return the bottom-up level order traversal of its nodes' values.
// (ie, from left to right, level by level from leaf to root).
//
//For example:
//Given binary tree [3,9,20,null,null,15,7],
//    3
//   / \
//  9  20
// /     \
//15      7
//return its bottom-up level order traversal as:
//[
//[15,7],
//[9,20],
//[3]
//]

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	var ret [][]int
	lastLevelNodes := []*TreeNode{root}
	for len(lastLevelNodes) != 0 {
		var levelArr []int
		var levelNodes []*TreeNode
		for _, t := range lastLevelNodes {
			levelArr = append(levelArr, t.Val)
			if t.Left != nil {
				levelNodes = append(levelNodes, t.Left)
			}
			if t.Right != nil {
				levelNodes = append(levelNodes, t.Right)
			}
		}
		lastLevelNodes = levelNodes
		result = append(result, levelArr)
	}
	for i := len(result) - 1; i >=0; i -- {
		ret = append(ret, result[i])
	}
	return ret
}


func main() {
	t4 := &TreeNode{1, nil, nil}
	t3 := &TreeNode{2, nil, nil}
	t2 := &TreeNode{3, t3, t4}
	fmt.Println(levelOrderBottom(t2))
}
