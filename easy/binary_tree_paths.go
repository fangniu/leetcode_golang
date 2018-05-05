package main

import (
	"strconv"
	"strings"
)

//Given a binary tree, return all root-to-leaf paths.
//
//For example, given the following binary tree:
//
//   1
// /   \
//2     3
// \
//  5
//All root-to-leaf paths are:
//
//["1->2->5", "1->3"]


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	val:=strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{val}
	}
	var str []string
	for _, value := range append(binaryTreePaths(root.Left), binaryTreePaths(root.Right)...) {
		str = append(str, val + "->" + value)
	}
	return str
}


func binaryTreePaths2(root *TreeNode) []string {
	var res []string
	if root != nil {
		search(root, "", &res)
	}
	return res
}

func search(root *TreeNode, path string, res *[]string) {
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path + strconv.Itoa(root.Val))
	}
	if root.Left != nil {
		search(root.Left, path + strconv.Itoa(root.Val) + "->", res)
	}
	if root.Right != nil {
		search(root.Right, path + strconv.Itoa(root.Val) + "->", res)
	}
}