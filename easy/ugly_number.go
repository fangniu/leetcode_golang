package main

import "fmt"

//Write a program to check whether a given number is an ugly number.
//
//Ugly numbers are positive numbers whose prime factors only include 2, 3, 5. For example, 6, 8 are ugly while 14 is not
// ugly since it includes another prime factor 7.
//
//Note:
//
//1 is typically treated as an ugly number.
//Input is within the 32-bit signed integer range.

func isUgly(num int) bool {
	arr := [3]int{2, 3, 5}
	for _, n := range arr {
		for num > 0 && num % n == 0 {
			num /= n
		}

	}
	return num == 1
}

func main() {
	fmt.Println(isUgly(7))
}