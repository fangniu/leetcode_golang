package main

import (
	"fmt"
)

//You are climbing a stair case. It takes n steps to reach to the top.
//
//Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
//Note: Given n will be a positive integer.
//
//Example 1:
//
//Input: 2
//Output: 2
//Explanation: There are two ways to climb to the top.
//1. 1 step + 1 step
//2. 2 steps
//Example 2:
//
//Input: 3
//Output: 3
//Explanation: There are three ways to climb to the top.
//1. 1 step + 1 step + 1 step
//2. 1 step + 2 steps
//3. 2 steps + 1 step

// 斐波那契的非递归实现
func climbStairs(n int) int {
	a := 1
	b := 1
	for i:=0; i < n; i++ {
		a, b = b, a + b
	}
	return a
}


// 递归效率太低
//func climbStairs2(n int) int {
//	if n <= 1 {
//		return 1
//	}
//	return climbStairs2(n-1) + climbStairs2(n-2)
//}



func main() {
	fmt.Println(climbStairs(20))
}