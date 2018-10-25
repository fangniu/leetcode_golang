package main

import (
	"strings"
	"fmt"
)

//
//Given two strings A and B, find the minimum number of times A has to be repeated such that B is a substring of it.
// If no such solution, return -1.
//
//For example, with A = "abcd" and B = "cdabcdab".
//
//Return 3, because by repeating A three times (“abcdabcdabcd”), B is a substring of it; and B is not a substring of A
// repeated two times ("abcdabcd").
//
//Note:
//The length of A and B will be between 1 and 10000.
//
//


func repeatedStringMatch(A string, B string) int {
	res := 0
	var newAS []int32
	for len(newAS) < len(B) {
		for _, s := range A {
			newAS = append(newAS, s)
		}
		res ++
	}
	newA := string(newAS)
	if strings.Index(newA, B) != -1 {
		return res
	}
	if strings.Index(newA+A, B) != -1 {
		return res + 1
	}
	return -1
}

func main() {
	fmt.Println(repeatedStringMatch("abcd", "cdabcdab"))
}