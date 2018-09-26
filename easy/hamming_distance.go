package main

import "fmt"

//The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
//
//Given two integers x and y, calculate the Hamming distance.
//
//Note:
//0 ≤ x, y < 231.
//
//Example:
//
//Input: x = 1, y = 4
//
//Output: 2
//
//Explanation:
//1   (0 0 0 1)
//4   (0 1 0 0)
//　　　　　↑   ↑
//
//The above arrows point to positions where the corresponding bits are different.


func hammingDistance(x int, y int) int {
	var count int
	z := x^y
	for z != 0 {
		if z&1 == 1 {
			count ++
		}
		z >>= 1
	}
	return count
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}