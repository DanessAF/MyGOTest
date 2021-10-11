package main
/*
 * @lc app=leetcode.cn id=538 lang=golang
 *
 * [538] 把二叉搜索树转换为累加树
 */

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Right)
			sum += root.Val
			root.Val = sum
			dfs(root.Left)
		}
	}
	dfs(root)
	return root
}

// @lc code=end

