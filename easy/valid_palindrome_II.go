package main

import "fmt"

//
//Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.
//
//Example 1:
//Input: "aba"
//Output: True
//Example 2:
//Input: "abca"
//Output: True
//Explanation: You could delete the character 'c'.
//Note:
//The string will only contain lowercase characters a-z. The maximum length of the string is 50000.


func validPalindrome(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j && s[i] == s[j]{
		i ++
		j --
	}
	if i >= j {
		return true
	}
	if isPalin(s, i+1, j) || isPalin(s, i, j-1) {
		return true
	}
	return false
}

func isPalin(s string, i int, j int) bool {
	for i < j {
		if s[i] == s[j] {
			i ++
			j --
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(validPalindrome("aba"))
}