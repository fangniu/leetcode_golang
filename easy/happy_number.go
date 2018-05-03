package main

import "fmt"

//Write an algorithm to determine if a number is "happy".
//
//A happy number is a number defined by the following process: Starting with any positive integer, replace the number
// by the sum of the squares of its digits, and repeat the process until the number equals 1 (where it will stay), or
// it loops endlessly in a cycle which does not include 1. Those numbers for which this process ends in 1 are happy
// numbers.
//
//Example:
//
//Input: 19
//Output: true
//Explanation:
//12 + 92 = 82
//82 + 22 = 68
//62 + 82 = 100
//12 + 02 + 02 = 1

func isHappy(n int) bool {
	slow := n
	fast := n
	slow = digitSquareSum(slow)
	fast = digitSquareSum(fast)
	fast = digitSquareSum(fast)
	for slow != fast {
		slow = digitSquareSum(slow)
		fast = digitSquareSum(fast)
		fast = digitSquareSum(fast)
	}
	fmt.Println(slow, fast)
	if slow == 1 {
		return true
	} else {
		return false
	}
}

func digitSquareSum(n int) int {
	var sum, tmp int
	for n > 0 {
		tmp = n % 10
		sum += tmp * tmp
		n /= 10
	}
	return sum
}

func main() {
	fmt.Println(isHappy(19))
}