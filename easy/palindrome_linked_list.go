package main

//Given a singly linked list, determine if it is a palindrome.
//
//Follow up:
//Could you do it in O(n) time and O(1) space?


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	var rev *ListNode
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		middle := slow.Next
		slow.Next = rev
		rev = slow
		slow = middle
	}
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if rev.Val != slow.Val {
			return false
		}
		rev, slow = rev.Next, slow.Next
	}
	return true
}