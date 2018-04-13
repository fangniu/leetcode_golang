package main

import "fmt"

//Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.
//
//You may assume no duplicates in the array.
//
//Example 1:
//
//Input: [1,3,5,6], 5
//Output: 2
//Example 2:
//
//Input: [1,3,5,6], 2
//Output: 1
//Example 3:
//
//Input: [1,3,5,6], 7
//Output: 4
//Example 1:
//
//Input: [1,3,5,6], 0
//Output: 0

func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || nums[0] >= target {
		return 0
	}
	if nums[len(nums) - 1] < target {
		return len(nums)
	}

	i := 0
	j := len(nums) - 1
	for i != j {
		if j - i == 1 {
			return j
		}
		middle := (i+j)/2
		if target == nums[middle]  {
			return  middle
		}
		if target < nums[middle] {
			j = middle
		} else {
			i = middle
		}
	}
	return i
}

func main() {
	fmt.Println(searchInsert([]int{1,3}, 1))
}
