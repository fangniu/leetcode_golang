package main


//Write a program to find the node at which the intersection of two singly linked lists begins.
//
//
//For example, the following two linked lists:
//
//A:          a1 → a2
//					↘
//					c1 → c2 → c3
//					↗
//B:     b1 → b2 → b3
//begin to intersect at node c1.
//
//
//Notes:
//
//If the two linked lists have no intersection at all, return null.
//The linked lists must retain their original structure after the function returns.
//You may assume there are no cycles anywhere in the entire linked structure.
//Your code should preferably run in O(n) time and use only O(1) memory.

func getIntersectionNode(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil && head2 == nil {
		return nil
	}
	a := head1
	b := head2
	for a != b {
		if a == nil {
			a = head2
		} else {
			a = a.Next
		}
		if b == nil {
			b = head1
		} else {
			b = b.Next
		}
	}
	return a
}