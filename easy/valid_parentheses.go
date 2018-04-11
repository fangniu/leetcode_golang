package main

import "fmt"

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
//The brackets must close in the correct order, "()" and "()[]{}" are all valid but "(]" and "([)]" are not.

func isValid(s string) bool {
	var stack []int32
	for _, b1 := range s {
		if len(stack) != 0 {
			b2 := stack[len(stack) - 1 ]
			if (b2 == '(' && b1 == ')') || (b2 == '{' && b1 == '}') || (b2 == '[' && b1 == ']') {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, b1)
			}
		} else if b1 == '(' || b1 == '{' || b1 == '[' {
			stack = append(stack, b1)
		} else {
			return false
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("()([])"))
}