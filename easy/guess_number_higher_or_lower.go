package main

import "math/rand"

//We are playing the Guess Game. The game is as follows:
//
//I pick a number from 1 to n. You have to guess which number I picked.
//
//Every time you guess wrong, I'll tell you whether the number is higher or lower.
//
//You call a pre-defined API guess(int num) which returns 3 possible results (-1, 1, or 0):
//
//-1 : My number is lower
//1 : My number is higher
//0 : Congrats! You got it!
//Example :
//
//Input: n = 10, pick = 6
//Output: 6


var result []int

func guessNumber(n int) int {
	for i := 0; i < n; i ++ {
		if guess(i) == 0 {
			return i
		}
	}
	return n
}

func guess(n int) int {
	return result[rand.Intn(3)]
}

func init() {
	result = []int{-1, 1, 0}
}