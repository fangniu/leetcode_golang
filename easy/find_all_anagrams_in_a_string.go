package main

import "fmt"

//
//Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.
//
//Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.
//
//The order of output does not matter.
//
//Example 1:
//
//Input:
//s: "cbaebabacd" p: "abc"
//
//Output:
//[0, 6]
//
//Explanation:
//The substring with start index = 0 is "cba", which is an anagram of "abc".
//The substring with start index = 6 is "bac", which is an anagram of "abc".
//Example 2:
//
//Input:
//s: "abab" p: "ab"
//
//Output:
//[0, 1, 2]
//
//Explanation:
//The substring with start index = 0 is "ab", which is an anagram of "ab".
//The substring with start index = 1 is "ba", which is an anagram of "ab".
//The substring with start index = 2 is "ab", which is an anagram of "ab".


func findAnagrams(s string, p string) []int {
	record := make([]int, 256)
	result := make([]int, 0)
	if len(s) == 0 || len(s) < len(p) {
		return result
	}
	for _, e := range p {
		record[e]++
	}
	left, right, count := 0, 0, len(p)
	for right < len(s) {
		if record[s[right]] >= 1 {
			count --
		}
		record[s[right]] --
		right ++
		if count == 0 {
			result = append(result, left)
		}
		if right - left == len(p) {
			if record[s[left]] >= 0 {
				count++
			}
			record[s[left]]++; left++
		}
	}
	return result
}

func main() {
	fmt.Println(findAnagrams("abc", "bc"))
}