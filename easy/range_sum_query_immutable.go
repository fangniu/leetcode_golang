package main

import "fmt"

//Given an integer array nums, find the sum of the elements between indices i and j (i ≤ j), inclusive.
//
//Example:
//Given nums = [-2, 0, 3, -5, 2, -1]
//
//sumRange(0, 2) -> 1
//sumRange(2, 5) -> -1
//sumRange(0, 5) -> -3
//Note:
//You may assume that the array does not change.
//There are many calls to sumRange function.


type NumArray struct {
	nums []int
	sums []int
}


func Constructor(nums []int) NumArray {
	var sums []int
	if len(nums) == 0 {
		return NumArray{}
	}
	sums = append(sums, nums[0])
	for i := 1; i < len(nums); i ++ {
		sums = append(sums, sums[i-1] + nums[i])
	}
	return NumArray{nums, sums}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i <= j && j < len(this.sums) {
		return this.sums[j] - this.sums[i] + this.nums[i]
	}
	return -1
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	obj := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(obj.SumRange(2,5))
}