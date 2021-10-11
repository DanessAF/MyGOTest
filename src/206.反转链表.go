package main

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

type ListNode struct {
	Val  int
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
func reverseList(head *ListNode) *ListNode {

}

func main() {
	testdata := &ListNode{1, nil}
	testdata.Next = &ListNode{2, nil}
	testdata.Next.Next = &ListNode{3, nil}
	testdata.Next.Next.Next = &ListNode{4, nil}
	testdata.Next.Next.Next.Next = &ListNode{5, nil}
	t := reverseList(testdata)
	for t != nil {
		println(t.Val)
		t = t.Next
	}
}

// @lc code=end
