package main

import "fmt"

//Find the nth digit of the infinite integer sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ...
//
//Note:
//n is positive and will fit within the range of a 32-bit signed integer (n < 231).
//
//Example 1:
//
//Input:
//3
//
//Output:
//3
//Example 2:
//
//Input:
//11
//
//Output:
//0
//
//Explanation:
//The 11th digit of the sequence 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, ... is a 0, which is part of the number 10.

// [1-9] [10-99] [100-999] [1000-9999] ...
func findNthDigit(n int) int {
	base, digit := 1, 1
	for {
		diff := n - (9 * base * digit)
		if diff < 0 {
			j := n % digit
			num := base + (n/digit - 1)
			if j != 0 {
				num++
			}
			if j == 0 {
				return num % 10
			}

			for ; j < digit; j++ {
				num /= 10
			}
			return num % 10
		}
		n, base, digit = diff, base*10, digit+1
	}
}


func main() {
	fmt.Println(findNthDigit(11000))
}