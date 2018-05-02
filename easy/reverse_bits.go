package main

import "fmt"

//Reverse bits of a given 32 bits unsigned integer.
//
//For example, given input 43261596 (represented in binary as 00000010100101000001111010011100),
// return 964176192 (represented in binary as 00111001011110000010100101000000).
//
//Follow up:
//If this function is called many times, how would you optimize it?
//
//Related problem: Reverse Integer

func reverseBits(n uint32) uint32 {
	if n == 0 {
		return 0
	}
	var result uint32
	for i := 0; i < 32; i ++ {
		result <<= 1
		if n & 1 == 1 {
			result ++
		}
		n >>= 1
	}
	return result
}

func main() {
	fmt.Println(reverseBits(43261596))
}