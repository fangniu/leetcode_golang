package main

import "fmt"

//Count the number of prime numbers less than a non-negative number, n.
//
//Example:
//
//Input: 10
//Output: 4
//Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.


func countPrimes(n int) int {
	if n <= 2 {
		return 0
	}
	notPrime := make([]bool, n)
	count := 1
	for i := 3; i < n; i += 2 {
		if notPrime[i] == false {
			count ++
			for j := 3; i*j < n; j += 2 {
				notPrime[i*j] = true
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(10))
}