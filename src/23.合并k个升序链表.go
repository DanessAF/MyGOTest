package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并K个升序链表
 */
type ListNode struct {
	Val int
	Next *ListNode
}
// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) (res *ListNode) {
	return
}

func main() {
	a := &ListNode{Val: 1}
	//b := &ListNode{Val: 1}
	//c := &ListNode{Val: 1}

	aa := []int{6, 5, 1}
	addList(a,0, aa)

	now := a
	for now.Next != nil {
		println(a.Val)
		now := now.Next
		println(now.Val)
	}
	fmt.Printf("%v", a)
	//mergeKLists([])
}

func addList(node *ListNode, idx int, nums []int)  {
	if idx > len(nums) - 1 {
		return
	}

	node.Val = nums[idx]
	node.Next = &ListNode{}
	addList(node.Next, idx + 1, nums)
}
// @lc code=end

