package main

import "fmt"

//Given an integer n, return the number of trailing zeroes in n!.
//
//Note: Your solution should be in logarithmic time complexity.

func trailingZeroes(n int) int {
	var count int
	for n > 0 {
		count += n/5
		n /= 5
	}
	return count
}

func main() {
	fmt.Println(trailingZeroes(85))
}