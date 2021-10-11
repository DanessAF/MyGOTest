package main

/*
 * @lc app=leetcode.cn id=27 lang=golang
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	i := 0
	for len(nums) > i {
		num := nums[i]
		if num == val {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
		i++
	}

	return len(nums)
}

// func main() {
// 	println(removeElement([]int{0,1,2,2,3,0,4,2},2))
// }
// @lc code=end

