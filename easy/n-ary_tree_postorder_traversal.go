package main

//Given an n-ary tree, return the postorder traversal of its nodes' values.
//
//
//For example, given a 3-ary tree:
//      1
//    / | \
//   3  2  4
//  / \
// 5   6
//Return its postorder traversal as: [5,6,3,2,4,1].
//
//
//Note: Recursive solution is trivial, could you do it iteratively?

type Node struct {
	Val int
	children []*Node
}

func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	for _, c := range root.children {
		for _, v := range postorder(c) {
			res = append(res, v)
		}
	}
	res = append(res, root.Val)
	return res
}