package main

import "fmt"

//Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.
//
//Note that the row index starts from 0.
//
//
//In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
//Example:
//
//Input: 3
//Output: [1,3,3,1]
//Follow up:
//
//Could you optimize your algorithm to use only O(k) extra space?

func getRow(rowIndex int) []int {
	if rowIndex == 0 {
		return []int{1}
	}
	last := getRow(rowIndex - 1)
	var ret []int
	for j := 0; j < rowIndex + 1; j ++ {
		if j == 0 || j == rowIndex {
			ret = append(ret, 1)
		} else {
			ret = append(ret, last[j-1] + last[j])
		}
	}
	return ret
}

func main() {
	fmt.Println(getRow(20))
}