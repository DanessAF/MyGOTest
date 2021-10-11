package main
/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */

// @lc code=start
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var ans []int
	ans = append(ans, 1)
	ans = append(ans, 2)
	for i := 2; i < n; i++ {
		ans = append(ans, ans[i - 2] + ans[i - 1])
	}
	return ans[n]
}

func main() {

}
// @lc code=end

