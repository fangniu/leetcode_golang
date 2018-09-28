package main

import (
	"fmt"
	"math"
)

//Given an integer array, find three numbers whose product is maximum and output the maximum product.
//
//Example 1:
//Input: [1,2,3]
//Output: 6
//Example 2:
//Input: [1,2,3,4]
//Output: 24
//Note:
//The length of the given array will be in range [3,104] and all elements are in the range [-1000, 1000].
//Multiplication of any three numbers in the input won't exceed the range of 32-bit signed integer.


//The idea is to track the top 3 elements and the bottom 2 elements as the max product could be achieved from top 3
//numbers or from the product of min negative numbers & max pos number.

func maximumProduct(a []int) int {
	top := []int{math.MinInt32, math.MinInt32, math.MinInt32}

	bot := []int{math.MaxInt32, math.MaxInt32}

	for i := 0; i < len(a); i++ {
		if a[i] > top[2] {
			top[0], top[1], top[2] = top[1], top[2], a[i]
		} else if a[i] > top[1] {
			top[0], top[1] = top[1], a[i]
		} else if a[i] > top[0] {
			top[0] = a[i]
		}

		if a[i] < bot[1] {
			bot[0], bot[1] = bot[1], a[i]
		} else if a[i] < bot[0] {
			bot[0] = a[i]
		}
	}

	res := top[0]*top[1]*top[2]

	if res < top[2]*bot[0]*bot[1] {
		return top[2]*bot[0]*bot[1]
	}

	return res
}

func main() {
	fmt.Println(maximumProduct([]int{1,2,3,4}))
}