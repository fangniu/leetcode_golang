package main

import "fmt"

//Given two binary strings, return their sum (also a binary string).
//
//The input strings are both non-empty and contains only characters 1 or 0.
//
//Example 1:
//
//Input: a = "11", b = "1"
//Output: "100"
//Example 2:
//
//Input: a = "1010", b = "1011"
//Output: "10101"

func addBinary(a string, b string) string {
	var s string
	carry := 0
	i := len(a) - 1
	j := len(b) - 1
	for i >=0 || j >=0 {
		sum := carry
		if i>= 0 && a[i] == '1' {
			sum ++
		}
		if j>= 0 && b[j] == '1' {
			sum ++
		}
		carry = sum / 2
		s = fmt.Sprintf("%v", sum % 2) + s
		i--
		j--
	}
	if carry > 0 {
		s = "1" + s
	}
	return s
}

func main() {
	fmt.Println(addBinary("11", "1"))
}