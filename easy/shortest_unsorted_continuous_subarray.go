package main

import (
	"fmt"
		)

//Given an integer array, you need to find one continuous subarray that if you only sort this subarray in ascending
// order, then the whole array will be sorted in ascending order, too.
//
//You need to find the shortest such subarray and output its length.
//
//Example 1:
//Input: [2, 6, 4, 8, 10, 9, 15]
//Output: 5
//Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.
//Note:
//Then length of the input array is in range [1, 10,000].
//The input array may contain duplicates, so ascending order here means <=.
//
//　　题目给了我们一个nums array， 让我们找出一个最短的无序连续子数组，当我们把这个子数组排序之后，整个array就已经是排序的了。
//
//　　要找到这个子数组的范围，先要了解这个范围的start 和 end 是如何定义的。
//
//　　来看这个例子：1 3 5 7 2 4 5 6
//
//　　a. 当我们找到第一个违反ascending 排序的数字 2的时候，我们不能是仅仅把start 标记为2的前面一个数字7，而是要一直往前，找到一个合适的位置，找到在
//      最前面位置的比2大的数字，这里是3。
//
//　　b. 同样的，为了找end， 那么我们要从7的后面开始找，一直找到一个最后面位置的比7小的数字，这里是6。
//
//　　这样的话，范围就是3到6 是我们要找的子数组。把3到6排序完了之后，整个array 就已经是排序的了。
//
//
//
//　　这里我们可以发现，2是min， 7是max，所以我们可以分两个方向来分别寻找start 和end。
//
//　　从右到左（绿色），维护更新min 和 start；
//
//　　从左到右（红色），维护更新max 和 end。



func findUnsortedSubarray(nums []int) int {
	length := len(nums)
	start := -1
	end := -2
	min := nums[length-1]
	max := nums[0]
	for i := 0; i < length; i ++ {
		max = maxInt(max, nums[i])
		min = minInt(min, nums[length - i -1])
		if nums[i] < max {
			end = i
		}
		if nums[length - i -1] > min {
			start = length - i -1
		}
	}
	return end - start + 1
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(findUnsortedSubarray([]int{1,2,5,4}))
}