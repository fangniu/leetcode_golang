package main

import "fmt"

//Given a 32-bit signed integer, reverse digits of an integer.
//
//Example 1:
//
//Input: 123
//Output:  321
//Example 2:
//
//Input: -123
//Output: -321
//Example 3:
//
//Input: 120
//Output: 21
//Note:
//Assume we are dealing with an environment which could only hold integers within the 32-bit signed integer range.
//For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

func reverse(x int) int {
	if x < 0 {
		return reverse(x * (-1)) * (-1)
	}
	var ret int
	for x != 0 {
		ret = ret * 10 + x % 10
		x = x / 10
	}
	if ret > (1<<31-1) {
		return 0
	}
	return ret
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-123))
	fmt.Println(reverse(120))
	fmt.Println(reverse(1563847412))
	fmt.Println(1<<31-1)

}