package main

import "math"

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	var min = math.MaxInt32
	var ans int
	for i, price := range prices {
		if price <= min {
			min = price
			continue
		}
		if ans < prices[i] - min {
			ans = prices[i] - min
		}
	}
	println(ans)
	println(min)
	return ans
}

func main() {
	println(maxProfit([]int{2,4,1}))
}
// @lc code=end
