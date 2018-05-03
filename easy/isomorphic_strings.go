package main

import "fmt"

//Given two strings s and t, determine if they are isomorphic.
//
//Two strings are isomorphic if the characters in s can be replaced to get t.
//
//All occurrences of a character must be replaced with another character while preserving the order of characters. No two characters may map to the same character but a character may map to itself.
//
//Example 1:
//
//Input: s = "egg", t = "add"
//Output: true
//Example 2:
//
//Input: s = "foo", t = "bar"
//Output: false
//Example 3:
//
//Input: s = "paper", t = "title"
//Output: true

func isIsomorphic(s string, t string) bool {
	ms := [256]byte{}
	mt := [256]byte{}
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		if ms[s[i]] == 0 && mt[t[i]] == 0 {
			ms[s[i]] = t[i]
			mt[t[i]] = s[i]
		} else {
			if ms[s[i]] != t[i] || mt[t[i]] != s[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("aa", "ab"))
}