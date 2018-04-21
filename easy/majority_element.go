package main

import "fmt"

//Given an array of size n, find the majority element. The majority element is the element that appears more than ⌊ n/2 ⌋ times.
//
//You may assume that the array is non-empty and the majority element always exist in the array.

func majorityElement(nums []int) int {
	var count int
	var data int
	for i := 0; i < len(nums); i ++ {
		if count == 0 {
			data = nums[i]
			count = 1
		} else {
			if data == nums[i] {
				count ++
			} else {
				count --
			}
		}
	}
	return data
}

func main() {
	fmt.Println(majorityElement([]int{2, 3, 2}))
}