package main

//Given a n-ary tree, find its maximum depth.
//
//The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.
//
//For example, given a 3-ary tree:
//    1
//   /| \
//  3 2  4
// /\
//5  6
//We should return its max depth, which is 3.
//
//Note:
//
//The depth of the tree is at most 1000.
//The total number of nodes is at most 5000.

type Node struct {
	Val int
	children []*Node
}


func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	m := 0
	for _, n := range root.children {
		m = max(maxDepth(n), m)
	}
	return m + 1
}