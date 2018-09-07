package main

import (
	"strings"
	"fmt"
)

//Given two non-negative integers num1 and num2 represented as string, return the sum of num1 and num2.
//
//Note:
//
//The length of both num1 and num2 is < 5100.
//Both num1 and num2 contains only digits 0-9.
//Both num1 and num2 does not contain any leading zero.
//You must not use any built-in BigInteger library or convert the inputs to integer directly.

func addStrings(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	max := 0
	if l1 < l2 {
		max = l2
	} else {
		max = l1
	}
	s := make([]string, max + 1)
	var carry uint8
	for i := 0; i < max; i ++ {
		var n1, n2 uint8
		if l1 - i > 0 {
			n1 = num1[l1 -1 - i]
		} else {
			n1 = 0 + '0'
		}

		if l2 - i > 0 {
			n2 = num2[l2 -1 - i]
		} else {
			n2 = 0 + '0'
		}
		sum := n1 + n2 - '0' * 2 + carry
		if sum >= 10 {
			s[max-i] = string(sum - 10 + '0')
			carry = 1
		} else {
			s[max-i] = string(sum + '0')
			carry = 0
		}
	}
	if carry > 0 {
		s[0] = "1"
	}
	return strings.Join(s, "")
}


func main() {
	fmt.Println(addStrings("1", "9"))
}