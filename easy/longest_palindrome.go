package main

import "fmt"

//Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.
//
//This is case sensitive, for example "Aa" is not considered a palindrome here.
//
//Note:
//Assume the length of given string will not exceed 1,010.
//
//Example:
//
//Input:
//"abccccdd"
//
//Output:
//7
//
//Explanation:
//One longest palindrome that can be built is "dccaccd", whose length is 7.

func longestPalindrome(s string) int {
	counts := make([]int, 52)
	for _, char := range s {
		var i int32
		if char >= 'a' {
			i = char - 'a' + 26
		} else {
			i = char - 'A'
		}
		counts[i] ++
	}
	single := false
	length := 0
	for _, c := range counts {
		if c == 0 {
			continue
		}
		if c == 1 {
			single = true
			continue
		}
		if c % 2 != 0 {
			single = true
			length += c -1
		} else {
			length += c
		}
	}
	if single {
		length ++
	}
	return length
}

func main() {
	fmt.Println(longestPalindrome("zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez"))
}