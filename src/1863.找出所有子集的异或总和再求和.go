package main

/*
 * @lc app=leetcode.cn id=1863 lang=golang
 *
 * [1863] 找出所有子集的异或总和再求和
 */

var ans [][]int
// @lc code=start
func subsetXORSum(nums []int) int {
	ans := 0
	len := len(nums)
	for i := 0; i < (1 << len); i++ {
		tmp := 0
		for j := 0; j < len; j++ {
			if (1 << j) & i == 0 {
				tmp ^= nums[j]
			}
		}
		ans += tmp
	}
	return ans
}

func main() {
	println(subsetXORSum([]int{5,1,6}))
}
// @lc code=end

