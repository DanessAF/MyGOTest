package main

import "fmt"

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

type ListNode struct {
	Val  int
	Next *ListNode
}
type Node struct {
	data int
	next *Node
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	var left []int
	ans := true
	for head != nil {
		left = append(left, head.Val)
		head = head.Next //移动指针
	}

	if len(left) == 1 {
		return true
	}

	for i := len(left) - 1; i >= 1; i -= 2 {
		if left[i] != left[0] {
			return false
		}
		left = left[1:i]
	}
	return ans
}
func Shownode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}

func main() {
	var head = new(ListNode)
	head.Val = 0
	var tail *ListNode
	tail = head //tail用于记录最末尾的结点的地址，刚开始tail的的指针指向头结点
	t := []int{1, 2, 1}
	for _, i := range t {
		var node = ListNode{Val: i}
		(*tail).Next = &node
		tail = &node
	}
	println(isPalindrome(head.Next)) //遍历结果
}

// @lc code=end
