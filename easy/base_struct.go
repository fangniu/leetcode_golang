package main

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) string() {
	ll := l
	var s []string
	for ll != nil {
		s = append(s, fmt.Sprintf("%v", ll.Val))
		ll = ll.Next
	}
	fmt.Println(strings.Join(s, "->"))
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func (t *TreeNode) preorderTraversal() {
	root := t
	fmt.Println(root.Val)
	if root.Left != nil {
		t.preorderTraversal()
	}
}

func preorderTraversal(root *TreeNode) {
	fmt.Print(root.Val, " ")
	if root.Left == nil && root.Right == nil {
		return
	}
	if root.Left == nil {
		fmt.Print("null", " ")
	} else {
		preorderTraversal(root.Left)
	}
	if root.Right == nil {
		fmt.Print("null", " ")
	} else {
		preorderTraversal(root.Right)
	}
}

//func main() {
//	t4 := &TreeNode{1, nil, nil}
//	//t3 := &TreeNode{2, nil, nil}
//	t2 := &TreeNode{3, nil, t4}
//	preorderTraversal(t2)
//}