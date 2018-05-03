package main


//given an array of integers and an integer k, find out whether there are two distinct indices i and j in the array such
// that nums[i] = nums[j] and the absolute difference between i and j is at most k.
//
//Example 1:
//
//Input: [1,2,3,1], k = 3
//Output: true
//Example 2:
//
//Input: [1,0,1,1], k = 1
//Output: true
//Example 3:
//
//Input: [1,2,1], k = 0
//Output: false

func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}
	mapping := map[int]int{}
	for i := 0; i < len(nums); i ++{
		n := nums[i]
		if v, ok := mapping[n]; ok {
			if i - v <= k {
				return true
			} else {
				mapping[n] = i
			}
		} else {
			mapping[n] = i
		}
	}
	return false
}