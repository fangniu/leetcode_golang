package main

import "fmt"

//
//Given a non-empty string check if it can be constructed by taking a substring of it and appending multiple copies of the substring together. You may assume the given string consists of lowercase English letters only and its length will not exceed 10000.
//
//
//
//Example 1:
//
//Input: "abab"
//Output: True
//Explanation: It's the substring "ab" twice.
//Example 2:
//
//Input: "aba"
//Output: False
//Example 3:
//
//Input: "abcabcabcabc"
//Output: True
//Explanation: It's the substring "abc" four times. (And the substring "abcabc" twice.)

// s的长度能被重复字符串的长度整除，再计算重复字符串的合法性即可

func repeatedSubstringPattern(s string) bool {
	l := len(s)
	for i := 1; i <= l/2; i ++ {
		m := l % i
		if m == 0 {
			if is_equal(s, i) {
				return true
			}
		}
	}
	return false
}

func is_equal(s string, step int) bool {
	l := len(s)
	for i := 0; i < step; i ++ {
		e := s[i]
		for j :=0; j < l; j += step {
			if e != s[i+j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(repeatedSubstringPattern("abcabcabcabca"))
}