package main

import (
	"fmt"
)

//给定一个整数数列，找出其中和为特定值的那两个数。
//
//你可以假设每个输入都只会有一种答案，同样的元素不能被重用。
//
//示例:
//
//给定 nums = [2, 7, 11, 15], target = 9
//
//因为 nums[0] + nums[1] = 2 + 7 = 9
//所以返回 [0, 1]


func twoSum(nums []int, target int) []int {
	l := len(nums)
	m := make(map[int]int, l)
	m[target - nums[0]] = 0
	for i := 1; i < l; i ++ {
		if value, ok := m[nums[i]]; ok {
			return []int{value, i}
		}
		m[target - nums[i]] = i
	}
	return nil
}

func main() {
	a := []int{3, 2, 4}
	fmt.Println(twoSum(a, 6))
}

