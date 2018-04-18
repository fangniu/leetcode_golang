package main

import "fmt"

//Implement int sqrt(int x).
//
//Compute and return the square root of x, where x is guaranteed to be a non-negative integer.
//
//Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.
//
//Example 1:
//
//Input: 4
//Output: 2
//Example 2:
//
//Input: 8
//Output: 2
//Explanation: The square root of 8 is 2.82842..., and since
//the decimal part is truncated, 2 is returned.

func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	i := 0
	j := x
	for i < j {
		middle := (i+j)/2
		if middle == i {
			return middle
		}
		tmp := middle * middle
		if tmp == x {
			return middle
		}
		if tmp < x {
			i = middle
		} else {
			j = middle
		}
	}
	return i
}

func main() {
	fmt.Println(mySqrt(16))
}