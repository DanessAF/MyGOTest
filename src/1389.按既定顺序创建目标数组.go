package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1389 lang=golang
 *
 * [1389] 按既定顺序创建目标数组
 */

// @lc code=start
func createTargetArray(nums []int, index []int) []int {
	targetArray := make([]int, len(nums))
	for i, _ := range index {
		copy(targetArray[index[i] + 1:], targetArray[index[i]:])
		targetArray[index[i]] = nums[i]
	}
	return targetArray
}

func main() {
	fmt.Printf("%v", createTargetArray([]int{0,1,2,3,4}, []int{0,1,2,2,1}))
}
// @lc code=end

