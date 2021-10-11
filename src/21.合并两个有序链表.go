package main

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {

	return l1
}

func main() {
	a := &ListNode{1, nil}
	a.Next = &ListNode{2, nil}
	a.Next.Next = &ListNode{4, nil}

	b := &ListNode{1, nil}
	b.Next = &ListNode{3, nil}
	b.Next.Next = &ListNode{4, nil}
	t := mergeTwoLists(a, b)
	for t != nil {
		println(t.Val)
		t = t.Next
	}
}

// @lc code=end
