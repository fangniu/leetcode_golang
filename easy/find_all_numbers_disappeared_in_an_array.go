package main


//Given an array of integers where 1 ≤ a[i] ≤ n (n = size of array), some elements appear twice and others appear once.
//
//Find all the elements of [1, n] inclusive that do not appear in this array.
//
//Could you do it without extra space and in O(n) runtime? You may assume the returned list does not count as extra space.
//
//Example:
//
//Input:
//[4,3,2,7,8,2,3,1]
//
//Output:
//[5,6]


func findDisappearedNumbers(nums []int) []int {
	i := 0
	for i < len(nums) {
		if nums[i] != i + 1 && nums[nums[i]-1] != nums[i]{
			tmp := nums[nums[i]-1]
			nums[nums[i]-1] = nums[i]
			nums[i] = tmp
		} else {
			i += 1
		}
	}
	res := make([]int, 0)
	for i := range nums {
		if nums[i] != i+1 {
			res = append(res, i+1)
		}
	}
	return res
}