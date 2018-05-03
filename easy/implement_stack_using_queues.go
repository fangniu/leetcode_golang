package main

import "fmt"

//Implement the following operations of a stack using queues.
//
//push(x) -- Push element x onto stack.
//pop() -- Removes the element on top of the stack.
//top() -- Get the top element.
//empty() -- Return whether the stack is empty.
//Notes:
//You must use only standard operations of a queue -- which means only push to back, peek/pop from front, size, and is empty operations are valid.
//Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.
//You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).

type MyStack struct {
	data []int
	position int
}


/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{nil, -1}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	if len(this.data) == 0 || len(this.data) == this.position + 1{
		this.data = append(this.data, x)
	} else {
		this.data[this.position + 1] = x
	}
	this.position ++
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.position != -1 {
		element := this.data[this.position]
		this.position --
		return element
	}
	return -1
}


/** Get the top element. */
func (this *MyStack) Top() int {
	if this.position != -1 {
		return this.data[this.position]
	}
	return -1
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.position < 0 {
		return true
	} else {
		return false
	}
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push(2)
	obj.Pop()
	obj.Push(3)
	fmt.Println(obj.data, obj.position)
}