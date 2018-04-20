package main

import (
	"unicode"
	"fmt"
)

//Given a string, determine if it is a palindrome, considering only alphanumeric characters and ignoring cases.
//
//Note: For the purpose of this problem, we define empty string as valid palindrome.
//
//Example 1:
//
//Input: "A man, a plan, a canal: Panama"
//Output: true
//Example 2:
//
//Input: "race a car"
//Output: false


func isPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		for i < j && !unicode.IsLetter(rune(s[i])) && !unicode.IsNumber(rune(s[i])) {
			i ++
		}
		for i < j && !unicode.IsLetter(rune(s[j])) && !unicode.IsNumber(rune(s[j])) {
			j --
		}
		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}
		i ++
		j --
	}
	return true
}

func main() {
	fmt.Println(isPalindrome("0p"))
}