package main

/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	var ans []int
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	ans = append(ans, nums[0])
	ans = append(ans, max(ans[0], nums[1]))

	for i := 2; i < len(nums); i++ {
		ans = append(ans, max(ans[i - 2] + nums[i], ans[i - 1]))
	}

	return ans[len(ans) - 1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end
