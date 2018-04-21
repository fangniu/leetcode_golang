package main

import "fmt"

//Given an array of integers, every element appears twice except for one. Find that single one.
//
//Note:
//Your algorithm should have a linear runtime complexity. Could you implement it without using extra memory?

func singleNumber(nums []int) int {
	var r int
	for _, n := range nums {
		r = r^n
	}
	return r
}

func main() {
	fmt.Println(singleNumber([]int{1, 1, 2, 3, 4, 4, 3}))
}