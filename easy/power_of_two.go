package main

import "fmt"

//Given an integer, write a function to determine if it is a power of two.

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	if n == 1 {
		return true
	}
	n --
	for n > 0 {
		if n & 1 != 1 {
			return false
		}
		n >>= 1
	}
	return true
}

func isPowerOfTwo2(n int) bool {
	if n == 0 {
		return false
	}

	return n&(n-1) == 0
}

func main() {
	fmt.Println(isPowerOfTwo(1023))
}