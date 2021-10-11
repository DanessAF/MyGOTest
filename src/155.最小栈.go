package main

import "math"

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	data []int
	min int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{min: math.MaxUint32}
}


func (this *MinStack) Push(val int)  {
	if val < this.min {
		this.min = val
	}
	this.data = append(this.data, val)
}


func (this *MinStack) Pop()  {
	if this.Top() == this.min {
		this.min = math.MaxUint32
	}
	this.data = this.data[:len(this.data) - 1]

	for _, datum := range this.data {
		if datum < this.min {
			this.min = datum
		}
	}
}


func (this *MinStack) Top() int {
	len := len(this.data)
	data := this.data
	return data[len - 1]
}


func (this *MinStack) GetMin() int {
	return this.min
}

func main() {
	d := Constructor()
	d.Push(-2)
	d.Push(0)
	d.Push(-3)
	for _, datum := range d.data {
		println(datum)
	}
	d.Pop()
	for _, datum := range d.data {
		println(datum)
	}
	d.Top()
	println(d.GetMin())
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

