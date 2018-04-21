package main

//Given a linked list, determine if it has a cycle in it.
//
//Follow up:
//Can you solve it without using extra space?


/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode(int x) : val(x), next(NULL) {}
 * };
 */
func hasCycle(root *ListNode) bool {
	flag := false
	defer func() {
		recover()
		flag = true
	}()
	if flag {
		return false
	}
	p := root.Next
	q := root
	for p != q {
		p = p.Next.Next
		q = q.Next
	}
	return true
}