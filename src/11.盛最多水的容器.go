package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	maxArea := 0

	i := 0
	j := len(height) - 1
	for i < j {
		if height[i] < height[j] {
			if  (j - i) * height[i] > maxArea {
				maxArea = (j - i) * height[i]
			}
			i++
		} else {
			if  (j - i) * height[j] > maxArea {
				maxArea = (j - i) * height[j]
			}

			j--
		}
	}
	return maxArea
}

func main() {
	println(maxArea([]int{8,20,1,2,3,4,5,6}))
}
// @lc code=end

