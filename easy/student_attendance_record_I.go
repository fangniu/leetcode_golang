package main

import "fmt"

//You are given a string representing an attendance record for a student. The record only contains the following three characters:
//'A' : Absent.
//'L' : Late.
//'P' : Present.
//A student could be rewarded if his attendance record doesn't contain more than one 'A' (absent) or more than two continuous 'L' (late).
//
//You need to return whether the student could be rewarded according to his attendance record.
//
//Example 1:
//Input: "PPALLP"
//Output: True
//Example 2:
//Input: "PPALLL"
//Output: False

func checkRecord(s string) bool {
	var countA int
	var countL int
	for _, c := range s {
		if c == 'L' {
			countL ++
		} else {
			countL = 0
		}
		if c == 'A' {
			countA ++
		}
		if countA > 1 || countL > 2 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(checkRecord("LALL"))
}