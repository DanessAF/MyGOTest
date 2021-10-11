package main

import "fmt"

/*
 * @lc app=leetcode.cn id=31 lang=golang
 *
 * [31] 下一个排列
 */

// @lc code=start
func nextPermutation(nums []int)  {
	n := len(nums)
	println(nums[n - 1], nums[n - 2])
	println(nums[n - 1] < nums[n - 2])
	for nums[n - 1] < nums[n - 2] {
		n--
	}

	for i := len(nums); i >= n; i-- {

	}
	println(n)
	fmt.Printf("%v", nums)
}

func main() {
	nextPermutation([]int{4,5,2,6,3,1})
}
// @lc code=end

