package main


//Given an n-ary tree, return the preorder traversal of its nodes' values.
//
//
//For example, given a 3-ary tree:
//      1
//    / | \
//   3  2  4
//  / \
// 5   6
//Return its preorder traversal as: [1,3,5,6,2,4].
//
//
//Note: Recursive solution is trivial, could you do it iteratively?
//
//
type Node struct {
	Val int
	children []*Node
}

func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	for _, c := range root.children {
		for _, v := range preorder(c) {
			res = append(res, v)
		}
	}
	return res
}