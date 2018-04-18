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
