package main

import (
	"fmt"
	"strings"
)

//The count-and-say sequence is the sequence of integers with the first five terms as following:
//
//1.     1
//2.     11
//3.     21
//4.     1211
//5.     111221
//1 is read off as "one 1" or 11.
//11 is read off as "two 1s" or 21.
//21 is read off as "one 2, then one 1" or 1211.
//Given an integer n, generate the nth term of the count-and-say sequence.
//
//Note: Each term of the sequence of integers will be represented as a string.
//
//Example 1:
//
//Input: 1
//Output: "1"
//Example 2:
//
//Input: 4
//Output: "1211"

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	lastStr := countAndSay(n - 1)
	var arr []string
	var b int32
	var count int

	for _, b2 := range lastStr {
		if count == 0 {
			b = b2
			count = 1
			continue
		}
		if b == b2 {
			count ++
		} else {
			arr = append(arr, fmt.Sprintf("%v", count), string(b))
			count = 1
			b = b2
		}
	}
	if count != 0 {
		arr = append(arr, fmt.Sprintf("%v", count), string(b))
	}
	return strings.Join(arr, "")
}

func main() {
	fmt.Println(countAndSay(4))
}