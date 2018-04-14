package main

import "fmt"

//Given a non-negative integer represented as a non-empty array of digits, plus one to the integer.
//
//You may assume the integer do not contain any leading zero, except the number 0 itself.
//
//The digits are stored such that the most significant digit is at the head of the list.


func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	var tmp []int
	var ret []int
	i := len(digits) - 1
	pre := 1
	for ; i >= 0; i-- {
		value := digits[i] + pre
		if value < 10 {
			tmp = append(tmp, value)
			break
		}
		tmp = append(tmp, 0)
		if i == 0 {
			tmp = append(tmp, 1)
		}
	}
	if i > 0 {
		for _, d := range digits[:i] {
			ret = append(ret, d)
		}
	}
	for j := len(tmp) - 1; j >= 0; j-- {
		ret = append(ret, tmp[j])
	}
	return ret
}

func main() {
	fmt.Println(plusOne([]int{9, 9, 9}))
}