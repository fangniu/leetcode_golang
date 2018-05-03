package main

//Reverse a singly linked list.
//
//Example:
//
//Input: 1->2->3->4->5->NULL
//Output: 5->4->3->2->1->NULL


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1 := head.Next
	p2 := head
	p2.Next = nil
	for p1 != nil {
		tmp := p1
		p1 = p1.Next
		tmp.Next = p2
		p2 = tmp
	}
	return p2
}

func main() {
	n1 := ListNode{1, nil}
	n2 := ListNode{2, &n1}
	n3 := ListNode{3, &n2}
	reverseList(&n3).string()
	reverseList(nil)
}