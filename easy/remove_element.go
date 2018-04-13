package main

import "fmt"

//Given an array and a value, remove all instances of that value in-place and return the new length.
//
//Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.
//
//The order of elements can be changed. It doesn't matter what you leave beyond the new length.
//
//Example:
//
//Given nums = [3,2,2,3], val = 3,
//
//Your function should return length = 2, with the first two elements of nums being 2.

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	for ; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		nums[j] = nums[i]
		j ++
	}
	fmt.Println(nums[0:j])
	return j
}

func main() {
	fmt.Println(removeElement([]int{3, 2, 2, 3}, 2))
}