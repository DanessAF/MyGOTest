package main
/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 丢失的数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	numLen := len(nums)
	count :=(1 + numLen ) * numLen / 2
	numSum := 0
	for _, num := range nums {
		numSum += num
	}
	return count - numSum
}

func main() {
	missingNumber([]int{1,12})
}
// @lc code=end

