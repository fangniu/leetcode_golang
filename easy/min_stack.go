package main

//Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
//
//push(x) -- Push element x onto stack.
//pop() -- Removes the element on top of the stack.
//top() -- Get the top element.
//getMin() -- Retrieve the minimum element in the stack.
//Example:
//MinStack minStack = new MinStack();
//minStack.push(-2);
//minStack.push(0);
//minStack.push(-3);
//minStack.getMin();   --> Returns -3.
//minStack.pop();
//minStack.top();      --> Returns 0.
//minStack.getMin();   --> Returns -2.

type MinStack struct {
	Stack    []int
	Pos      int
	MinStack []int
	MinPos   int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack:    make([]int, 10240),
		MinStack: make([]int, 10240),
		Pos:      -1,
		MinPos:   -1,
	}
}


func (this *MinStack) Push(x int)  {
	this.Pos ++
	this.Stack[this.Pos] = x
	if this.MinPos == -1 || x <= this.MinStack[this.MinPos] {
		this.MinPos ++
		this.MinStack[this.MinPos] = x
	}
}


func (this *MinStack) Pop()  {
	if this.Pos == -1 {
		return
	}
	x := this.Stack[this.Pos]
	if x == this.MinStack[this.MinPos] {
		this.MinPos --
	}
	this.Pos --
}


func (this *MinStack) Top() int {
	if this.Pos == -1 {
		return -1
	}
	return this.Stack[this.Pos]
}


func (this *MinStack) GetMin() int {
	if this.Pos == -1 {
		return -1
	}
	return this.MinStack[this.MinPos]
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */