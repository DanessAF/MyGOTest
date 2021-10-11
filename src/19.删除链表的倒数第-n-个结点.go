package main

import "fmt"

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	
}



func main()  {
	testdata := &ListNode{1, nil}
	//testdata.Next = &ListNode{2, nil}
	//testdata.Next.Next = &ListNode{3, nil}
	//testdata.Next.Next.Next = &ListNode{4, nil}
	//testdata.Next.Next.Next.Next = &ListNode{5, nil}

	removeNthFromEnd(testdata, 1)

	fmt.Println("--------done----------")
}
// @lc code=end

