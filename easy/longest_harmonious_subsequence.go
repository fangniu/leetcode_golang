package main

import "fmt"

//We define a harmonious array is an array where the difference between its maximum value and its minimum value is exactly 1.
//
//Now, given an integer array, you need to find the length of its longest harmonious subsequence among all its possible subsequences.
//
//Example 1:
//Input: [1,3,2,2,5,2,3,7]
//Output: 5
//Explanation: The longest harmonious subsequence is [3,2,2,2,3].
//Note: The length of the input array will not exceed 20,000.
//


func findLHS(nums []int) int {
	m := map[int]int{}
	var maxCount int
	for _, n := range nums {
		if v, ok := m[n]; ok {
			m[n] = v + 1
		} else {
			m[n] = 1
		}

	}
	for k, v := range m {
		if c, ok := m[k+1]; ok {
			if v + c > maxCount {
				maxCount = v + c
			}
		}
	}
	return maxCount
}

func main() {
	fmt.Println(findLHS([]int{1,3,5,7,9,11,13,15,17}))
}