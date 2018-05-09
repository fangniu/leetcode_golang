package main


//Write a function that takes a string as input and returns the string reversed.
//
//Example:
//Given s = "hello", return "olleh".


func reverseString(s string) string {
	slen := len(s)
	if slen == 0 {
		return ""
	}
	left, right := 0, slen - 1
	sb := []byte(s)
	for left < right {
		sb[left], sb[right] = sb[right], sb[left]
		left, right = left + 1, right - 1
	}
	return string(sb)
}