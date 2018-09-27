package main

import "fmt"

//Given an integer array with even length, where different numbers in this array represent different kinds of candies.
// Each number means one candy of the corresponding kind. You need to distribute these candies equally in number to
// brother and sister. Return the maximum number of kinds of candies the sister could gain.
//Example 1:
//Input: candies = [1,1,2,2,3,3]
//Output: 3
//Explanation:
//There are three different kinds of candies (1, 2 and 3), and two candies for each kind.
//Optimal distribution: The sister has candies [1,2,3] and the brother has candies [1,2,3], too.
//The sister has three different kinds of candies.
//Example 2:
//Input: candies = [1,1,2,3]
//Output: 2
//Explanation: For example, the sister has candies [2,3] and the brother has candies [1,1].
//The sister has two different kinds of candies, the brother has only one kind of candies.
//Note:
//
//The length of the given array is in range [2, 10,000], and will be even.
//The number in given array is in range [-100,000, 100,000].

func distributeCandies(candies []int) int {
	unique := make([]int, 200001, 200001)
	var numUnique int
	for i := 0; i < len(candies); i++ {
		c := candies[i] + 100000
		was := unique[c]
		if was == 0 {
			numUnique++
		}
		unique[c] = was + 1
	}
	get := len(candies)/2
	if get < numUnique {
		return get
	}
	return numUnique
}

func main() {
	fmt.Println(distributeCandies([]int{1,1,2,2,3,3}))
}