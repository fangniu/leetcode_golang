package main

import "fmt"

//Given two strings s and t, write a function to determine if t is an anagram of s.
//
//For example,
//s = "anagram", t = "nagaram", return true.
//s = "rat", t = "car", return false.
//
//Note:
//You may assume the string contains only lowercase alphabets.
//
//Follow up:
//What if the inputs contain unicode characters? How would you adapt your solution to such case?

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapping := make(map[byte]int, 26)
	for i := 0; i < len(s); i++ {
		c1 := s[i]
		c2 := t[i]
		if c1 == c2 {
			continue
		}
		v1, ok := mapping[c1]
		if ok {
			mapping[c1] = v1 + 1
		} else {
			mapping[c1] = 1
		}
		v2, ok := mapping[c2]
		if ok {
			mapping[c2] = v2 - 1
		} else {
			mapping[c2] = -1
		}
	}
	for _, v := range mapping {
		if v != 0 {
			return false
		}
	}
	return true
}


func isAnagram2(s string, t string) bool {
	r := make([]int, 26)

	for _, v := range s {
		r[v-'a']++
	}

	for _, v := range t {
		r[v-'a']--
	}

	for i := 0; i < len(r); i++ {
		if r[i] != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}