package main

import "fmt"

//Given a string s consists of upper/lower-case alphabets and empty space characters ' ', return the length of last word in the string.
//
//If the last word does not exist, return 0.
//
//Note: A word is defined as a character sequence consists of non-space characters only.
//
//Example:
//
//Input: "Hello World"
//Output: 5

func lengthOfLastWord(s string) int {
	if len(s) == 0 {
		return 0
	}
	var count int
	for i := len(s) - 1; i >= 0; i -- {
		if s[i] == ' ' && count == 0 {
			continue
		}
		if s[i] == ' ' {
			break
		}
		count ++
	}
	return count
}

func main() {
	fmt.Println(lengthOfLastWord("  "))
}