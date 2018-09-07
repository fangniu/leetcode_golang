package main

import "fmt"

//Given a string, find the first non-repeating character in it and return it's index. If it doesn't exist, return -1.
//
//Examples:
//
//s = "leetcode"
//return 0.
//
//s = "loveleetcode",
//return 2.
//Note: You may assume the string contain only lowercase letters.

func firstUniqChar(s string) int {
	arr := [26]int{}
	for _, e := range s {
		arr[e - 'a'] ++
	}
	for i, e := range s {
		if arr[e - 'a'] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))
}