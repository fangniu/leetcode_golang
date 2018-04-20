package main

import "fmt"

//Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
//In Pascal's triangle, each number is the sum of the two numbers directly above it.
//
//Example:
//
//Input: 5
//Output:
//[
//[1],
//[1,1],
//[1,2,1],
//[1,3,3,1],
//[1,4,6,4,1]
//]

func generate(numRows int) [][]int {
	var s [][]int
	for i:=0; i < numRows; i ++ {
		var tmp []int
		for j := 0; j < i + 1; j ++ {
			if j == 0 || j == i {
				tmp = append(tmp, 1)
			} else {
				tmp = append(tmp, s[i-1][j-1] + s[i-1][j])
			}
		}
		s = append(s, tmp)
	}
	return s
}

func main() {
	fmt.Println(generate(0))
}