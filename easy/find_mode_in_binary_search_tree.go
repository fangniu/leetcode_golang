package main

import "fmt"

//Given a binary search tree (BST) with duplicates, find all the mode(s) (the most frequently occurred element) in the given BST.
//
//Assume a BST is defined as follows:
//
//The left subtree of a node contains only nodes with keys less than or equal to the node's key.
//The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
//Both the left and right subtrees must also be binary search trees.
//
//
//For example:
//Given BST [1,null,2,2],
//
//1
// \
//  2
// /
//2

//return [2].
//
//Note: If a tree has more than one mode, you can return them in any order.
//
//Follow up: Could you do that without using any extra space? (Assume that the implicit stack space incurred due to recursion does not count).
//

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


type mySolution struct {
	prev *int
	count int
	maxCount int
	list []int
}

func newSolution() *mySolution{
	s := mySolution{count:1}
	return &s
}

func (s *mySolution) traverse(root *TreeNode) {
	if root == nil {
		return
	}
	s.traverse(root.Left)
	if s.prev != nil {
		if root.Val == *s.prev {
			s.count ++
		} else {
			s.count = 1
		}
	}
	if s.count > s.maxCount {
		s.maxCount = s.count
		s.list = s.list[:0]
		s.list = append(s.list, root.Val)
	} else if s.count == s.maxCount {
		s.list = append(s.list, root.Val)
	}
	s.prev = &root.Val
	s.traverse(root.Right)
}

func findMode(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	s := newSolution()
	s.traverse(root)
	return s.list
}



func main() {
	n := TreeNode{2147483647, nil, nil}
	fmt.Println(findMode(&n))
}