package main

import "fmt"

//Given a positive integer, return its corresponding column title as appear in an Excel sheet.
//
//For example:
//
//1 -> A
//2 -> B
//3 -> C
//...
//26 -> Z
//27 -> AA
//28 -> AB

func convertToTitle(n int) string {
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var s []uint8
	for n > 0 {
		s = append(s, chars[(n - 1)%26])
		n = (n - 1) / 26
	}
	for i := 0; i < len(s)/2; i++ {
		j := len(s) - i - 1
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func main() {
	fmt.Println(convertToTitle(28))
}