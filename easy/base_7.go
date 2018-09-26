package main

import (
			"fmt"
	"strconv"
)

//Given an integer, return its base 7 string representation.
//
//Example 1:
//Input: 100
//Output: "202"
//Example 2:
//Input: -7
//Output: "-10"
//Note: The input will be in range of [-1e7, 1e7].


func convertToBase7(num int) string {
	if num < 0 {
		return "-" + convertToBase7(num * -1)
	}
	if num < 7 {
		return strconv.Itoa(num)
	}
	return convertToBase7(num / 7) + strconv.Itoa(num % 7)
}

func main() {
	fmt.Println(convertToBase7(100))
}