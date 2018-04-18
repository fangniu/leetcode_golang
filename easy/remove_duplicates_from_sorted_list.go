package main


//Given a sorted linked list, delete all duplicates such that each element appear only once.
//
//For example,
//Given 1->1->2, return 1->2.
//Given 1->1->1->2->3->3, return 1->2->3.


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	cur := head
	for cur != nil {
		tmp := cur
		for tmp.Next != nil && cur.Val == tmp.Next.Val {
			tmp = tmp.Next
		}
		cur.Next = tmp.Next
		cur = cur.Next
	}
	return head
}

func main() {
	l14 := &ListNode{4, nil}
	l13 := &ListNode{1, l14}
	l12 := &ListNode{1, l13}
	l11 := &ListNode{1, l12}
	deleteDuplicates(l11).string()
}