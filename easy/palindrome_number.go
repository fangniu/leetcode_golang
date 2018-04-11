package main

import "fmt"

//Determine whether an integer is a palindrome. Do this without extra space.
//
//click to show spoilers.
//
//Some hints:
//Could negative integers be palindromes? (ie, -1)
//
//If you are thinking of converting the integer to string, note the restriction of using extra space.
//
//You could also try reversing an integer. However, if you have solved the problem "Reverse Integer", you know that the reversed integer might overflow. How would you handle such case?
//
//There is a more generic way of solving this problem.

func isPalindrome(x int) bool {
	if x<0 || (x!=0 && x%10==0) {
		return false
	}
	rev := 0
	for x > rev {
		rev = rev * 10 + x % 10
		x = x / 10
	}
	if x == rev {
		return true
	}
	return x == rev/10
}

func main() {
	fmt.Println(isPalindrome(121))
}