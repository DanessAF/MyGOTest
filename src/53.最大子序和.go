package main

import "math"

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子序和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	ans := math.MinInt32
	tmp := 0
	for _, num := range nums {
		if tmp < 0 {
			tmp = num
		} else {
			tmp += num
		}
		if tmp > ans {
			 ans = tmp
		}
	}

	println(ans)
	return ans
}

func main() {
	maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
}

// @lc code=end
