package main

import "fmt"

//Given a word, you need to judge whether the usage of capitals in it is right or not.
//
//We define the usage of capitals in a word to be right when one of the following cases holds:
//
//All letters in this word are capitals, like "USA".
//All letters in this word are not capitals, like "leetcode".
//Only the first letter in this word is capital if it has more than one letter, like "Google".
//Otherwise, we define that this word doesn't use capitals in a right way.
//Example 1:
//Input: "USA"
//Output: True
//Example 2:
//Input: "FlaG"
//Output: False
//Note: The input will be a non-empty word consisting of uppercase and lowercase latin letters.


func detectCapitalUse(word string) bool {
	if len(word) == 1 {
		return true
	}
	firstCapital := false
	secondCapital := false
	if word[0] >= 'A' && word[0] <= 'Z' {
		firstCapital = true
	}

	if word[1] >= 'A' && word[1] <= 'Z' {
		secondCapital = true
	}

	if firstCapital && secondCapital {
		for _, w := range word[2:] {
			if w < 'A' || w > 'Z' {
				return false
			}
		}
		return true
	} else if firstCapital {
		for _, w := range word[2:] {
			if w < 'a' || w > 'z' {
				return false
			}
		}
		return true
	} else if secondCapital{
		return false
	} else {
		for _, w := range word[2:] {
			if w < 'a' || w > 'z' {
				return false
			}
		}
		return true
	}
}

func main() {
	fmt.Println(detectCapitalUse("FlaG"))
}