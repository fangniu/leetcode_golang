package main

//Given an array nums, write a function to move all 0's to the end of it while maintaining the relative order of the non-zero elements.
//
//For example, given nums = [0, 1, 0, 3, 12], after calling your function, nums should be [1, 3, 12, 0, 0].
//
//Note:
//You must do this in-place without making a copy of the array.
//Minimize the total number of operations.



func moveZeroes(nums []int)  {
	if len(nums) == 0 {
		return
	}
	insertIndex := 0
	for _, e := range nums {
		if e != 0 {
			nums[insertIndex] = e
			insertIndex ++
		}
	}
	for ; insertIndex < len(nums); insertIndex ++ {
		nums[insertIndex] = 0
	}
}
