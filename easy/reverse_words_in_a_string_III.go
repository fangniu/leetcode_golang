package main

import (
	"strings"
	"fmt"
)

//Given a string, you need to reverse the order of characters in each word within a sentence while still preserving
//whitespace and initial word order.
//
//Example 1:
//Input: "Let's take LeetCode contest"
//Output: "s'teL ekat edoCteeL tsetnoc"
//Note: In the string, each word is separated by single space and there will not be any extra space in the string.
//


func reverseWords(s string) string {
	var res []string
	for _, l := range strings.Split(s, " ") {
		newS := []byte(l)
		i := 0
		j := len(l) - 1
		for i < j {
			newS[i], newS[j] = l[j], l[i]
			i ++
			j --
		}
		res = append(res, string(newS))
	}
	return strings.Join(res, " ")
}

func main() {
	fmt.Println(reverseWords("Let's take LeetCode contest"))
}