package main

import (
	"fmt"
	"container/heap"
)

//
//Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted
// order, not the kth distinct element.
//
//Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains
// initial elements from the stream. For each call to the method KthLargest.add, return the element representing the
// kth largest element in the stream.
//
//Example:
//
//int k = 3;
//int[] arr = [4,5,8,2];
//KthLargest kthLargest = new KthLargest(3, arr);
//kthLargest.add(3);   // returns 4
//kthLargest.add(5);   // returns 5
//kthLargest.add(10);  // returns 5
//kthLargest.add(9);   // returns 8
//kthLargest.add(4);   // returns 8
//Note:
//You may assume that nums' length ≥ k-1 and k ≥ 1.

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

type KthLargest struct {
	Nums *IntHeap
	K    int
}


func Constructor(k int, nums []int) KthLargest {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < len(nums); i ++ {
		heap.Push(h, nums[i])
	}
	for len(*h) > k {
		heap.Pop(h)
	}
	return KthLargest{h, k}
}


func (this *KthLargest) Add(val int) int {
	if len(*this.Nums) < this.K {
		heap.Push(this.Nums, val)
	} else if val > (*this.Nums)[0] {
		heap.Push(this.Nums, val)
		heap.Pop(this.Nums)
	}
	return (*this.Nums)[0]
}


/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

func main() {
	k := 3
	arr := []int{4,5,8,2}
	obj := Constructor(k, arr)
	fmt.Println(obj.Add(3))
	fmt.Println(obj.Add(5))
	fmt.Println(obj.Add(10))
	fmt.Println(obj.Add(9))
	fmt.Println(obj.Add(4))
}