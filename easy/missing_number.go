package main

import "fmt"

//Given an array containing n distinct numbers taken from 0, 1, 2, ..., n, find the one that is missing from the array.
//
//Example 1
//
//Input: [3,0,1]
//Output: 2
//Example 2
//
//Input: [9,6,4,2,3,5,7,0,1]
//Output: 8
//
//Note:
//Your algorithm should run in linear runtime complexity. Could you implement it using only constant extra space complexity?


func missingNumber(nums []int) int {
	var sum int
	for i := 0; i < len(nums); i ++ {
		sum += nums[i]
	}
	return len(nums) * (len(nums) + 1) / 2  - sum
}

func missingNumber2(nums []int) int {
	var xor, i int
	for i = 0; i < len(nums); i ++ {
		xor = xor ^ i ^ nums[i]
	}
	return xor ^ i
}

func main() {
	fmt.Println(missingNumber2([]int{9,6,4,2,3,5,7,0,1}))
}
