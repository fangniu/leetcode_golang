package main

import (
	"fmt"
	"strings"
)

//Merge two sorted linked lists and return it as a new list.
//The new list should be made by splicing together the nodes of the first two lists.
//
//Example:
//
//Input: 1->2->4, 1->3->4
//Output: 1->1->2->3->4->4

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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


func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	tail := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		tail = tail.Next
	}
	if l1 == nil {
		tail.Next = l2
	} else {
		tail.Next = l1
	}
	return head.Next
}

func main() {
	l13 := ListNode{4, nil}
	l12 := ListNode{2, &l13}
	l11 := ListNode{1, &l12}

	l11.string()
	l23 := ListNode{4, nil}
	l22 := ListNode{3, &l23}
	l21 := ListNode{1, &l22}
	l21.string()
	l := mergeTwoLists(&l11, &l21)
	l.string()

}