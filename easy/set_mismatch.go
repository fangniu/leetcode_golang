package main

import "fmt"

//The set S originally contains numbers from 1 to n. But unfortunately, due to the data error, one of the numbers in the
// set got duplicated to another number in the set, which results in repetition of one number and loss of another number.
//
//Given an array nums representing the data status of this set after the error. Your task is to firstly find the number
// occurs twice and then find the number that is missing. Return them in the form of an array.
//
//Example 1:
//Input: nums = [1,2,2,4]
//Output: [2,3]
//Note:
//The given array size will in the range [2, 10000].
//The given array's numbers won't have any order.

func findErrorNums(nums []int) []int {
	freq := make([]int, len(nums) + 1)
	var num int
	var sum int
	for _, n:= range nums {
		sum += n
		if freq[n] == 1 {
			num = n
		} else {
			freq[n] ++
		}
	}
	return []int{num, (1+len(nums))*len(nums)/2-sum+num}
}

func main() {
	fmt.Println(findErrorNums([]int{1,2,2,4}))
}