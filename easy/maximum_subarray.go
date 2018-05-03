package main

import (
	"fmt"
)

//Find the contiguous subarray within an array (containing at least one number) which has the largest sum.
//
//For example, given the array [-2,1,-3,4,-1,2,1,-5,4],
//the contiguous subarray [4,-1,2,1] has the largest sum = 6.
//
//click to show more practice.
//
//More practice:
//If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.


func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curSum := nums[0]
	maxSum := curSum
	for _, i := range nums[1:] {
		curSum = max(i, i + curSum)
		maxSum = max(maxSum, curSum)
	}
	return int(maxSum)
}

func main() {
	fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}))
}