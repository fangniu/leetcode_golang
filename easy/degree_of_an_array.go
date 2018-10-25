package main

import "fmt"

//Given a non-empty array of non-negative integers nums, the degree of this array is defined as the maximum frequency of
// any one of its elements.
//
//Your task is to find the smallest possible length of a (contiguous) subarray of nums, that has the same degree as nums.
//
//Example 1:
//Input: [1, 2, 2, 3, 1]
//Output: 2
//Explanation:
//The input array has a degree of 2 because both elements 1 and 2 appear twice.
//Of the subarrays that have the same degree:
//[1, 2, 2, 3, 1], [1, 2, 2, 3], [2, 2, 3, 1], [1, 2, 2], [2, 2, 3], [2, 2]
//The shortest length is 2. So return 2.
//Example 2:
//Input: [1,2,2,3,1,4,2]
//Output: 6
//Note:
//
//nums.length will be between 1 and 50,000.
//nums[i] will be an integer between 0 and 49,999.

func findShortestSubArray(nums []int) int {
	left := make(map[int]int)
	right := make(map[int]int)
	count := make(map[int]int)
	degree := 0

	for i := 0; i < len(nums); i ++ {
		n := nums[i]
		if _, ok := left[n]; !ok {
			left[n] = i
		}
		right[n] = i
		if v, ok := count[n]; ok {
			count[n] = v + 1
		} else {
			count[n] = 1
		}
		if count[n] > degree {
			degree = count[n]
		}
	}
	ans := len(nums)
	for k, v := range count {
		if v == degree {
			length := right[k] - left[k] + 1
			if length < ans {
				ans = length
			}
		}
	}
	return ans
}

func main() {
	fmt.Println(findShortestSubArray([]int{1, 2, 2, 3, 1}))
}