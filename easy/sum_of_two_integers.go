package main

import "fmt"

//Calculate the sum of two integers a and b, but you are not allowed to use the operator + and -.
//
//Example:
//Given a = 1 and b = 2, return 3.

func getSum(a int, b int) int {
	for b != 0 {
		carry := a & b
		a = a ^ b
		b = carry << 1
	}
	return a
}

func main() {
	fmt.Println(getSum(11, 22))
}