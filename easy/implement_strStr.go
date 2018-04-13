package main

import "fmt"

//Implement strStr().
//
//Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
//
//Example 1:
//
//Input: haystack = "hello", needle = "ll"
//Output: 2
//Example 2:
//
//Input: haystack = "aaaaa", needle = "bba"
//Output: -1

func strStr(haystack string, needle string) int {
	if len(haystack) < len(needle) {
		return -1
	}
	for i := 0; i < len(haystack)-len(needle) + 1; i++ {
		if haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("mississippi", "pi"))
	fmt.Println(strStr("hello", "ll"))
	fmt.Println(strStr("", ""))
	fmt.Println(strStr("mississippi", "issi"))
}