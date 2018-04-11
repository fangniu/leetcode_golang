package main

import "fmt"

// Write a function to find the longest common prefix string amongst an array of strings.

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 || len(strs[0]) == 0 {
		return ""
	}
	var common []byte
	minLen := 0
	for _, str := range strs {
		if len(str) == 0 {
			return ""
		}
		if minLen == 0 {
			minLen = len(str)
		} else if len(str) < minLen {
			minLen = len(str)
		}
	}

	for i := 0; i < minLen; i++ {
		tmp := strs[0][i]
		for j := 1; j < len(strs); j ++ {
			if strs[j][i] != tmp {
				return string(common)
			}
		}
		common = append(common, tmp)
	}
	return string(common)
}

func main() {
	ret := longestCommonPrefix([]string{"abc", "abx", "ar"})
	fmt.Println(ret)
}