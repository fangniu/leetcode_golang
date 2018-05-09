package main

import "fmt"

//Given a positive integer num, write a function which returns True if num is a perfect square else False.
//
//Note: Do not use any built-in library function such as sqrt.
//
//Example 1:
//
//Input: 16
//Returns: True
//Example 2:
//
//Input: 14
//Returns: False

func isPerfectSquare(num int) bool {
	low := 0
	high := num
	for low <= high {
		m := (low + high) >> 1
		tmp := m * m
		if tmp == num {
			return true
		}
		if tmp < num {
			low = m + 1
		} else {
			high = m - 1
		}
	}
	return false
}

func main() {
	fmt.Println(isPerfectSquare(16))
}