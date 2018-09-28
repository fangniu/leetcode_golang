package main

import (
	"fmt"
	"math"
)

//Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.
//
//Example 1:
//Input: 5
//Output: True
//Explanation: 1 * 1 + 2 * 2 = 5
//Example 2:
//Input: 3
//Output: False


func judgeSquareSum(c int) bool {
	i := 0
	j := int(math.Sqrt(float64(c)))
	for i <= j {
		sum := i*i + j*j
		if sum == c {
			return true
		}
		if sum < c {
			i ++
		} else {
			j --
		}
	}
	return false
}

func main() {
	fmt.Println(judgeSquareSum(8))
}