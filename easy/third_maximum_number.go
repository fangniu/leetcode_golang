package main

import (
	"fmt"
	"math"
)

//Given a non-empty array of integers, return the third maximum number in this array. If it does not exist, return the maximum number. The time complexity must be in O(n).
//
//Example 1:
//Input: [3, 2, 1]
//
//Output: 1
//
//Explanation: The third maximum is 1.
//Example 2:
//Input: [1, 2]
//
//Output: 2
//
//Explanation: The third maximum does not exist, so the maximum (2) is returned instead.
//Example 3:
//Input: [2, 2, 3, 1]
//
//Output: 1
//
//Explanation: Note that the third maximum here means the third maximum distinct number.
//Both numbers with value 2 are both considered as second maximum.
//

func thirdMax(nums []int) int {
	min := int(math.Inf(-1))
	save := [] int {min,min,min}
	for _,v := range nums {
		if v != save[0] && v != save[1] && v != save[2]{
			if v > save[0]{
				save = [] int {v,save[0],save[1]}
			}else if v > save[1]{
				save = [] int {save[0],v,save[1]}
			}else if v > save[2]{
				save = [] int {save[0],save[1],v}
			}
		}
	}
	if save[2] == min{
		return save[0]
	}
	return save[2]
}

func main() {
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}