package main

import (
			"fmt"
)

//Given a string and an integer k, you need to reverse the first k characters for every 2k characters counting from the
//start of the string. If there are less than k characters left, reverse all of them. If there are less than 2k but
//greater than or equal to k characters, then reverse the first k characters and left the other as original.
//Example:
//Input: s = "abcdefg", k = 2
//Output: "bacdfeg"
//Restrictions:
//The string consists of lower English letters only.
//Length of the given string and k will in the range [1, 10000]

func reverseStr(s string, k int) string {
	res := []byte(s)
	for start := 0; start < len(s); start += 2*k {
		i := start
		j := min(start + k - 1, len(s) - 1)
		for i < j {
			res[i], res[j] = s[j], s[i]
			i ++
			j --
		}
	}
	return string(res)
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	fmt.Println(reverseStr("abcdefg", 2))
}