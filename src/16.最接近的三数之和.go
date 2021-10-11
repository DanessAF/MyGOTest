package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=16 lang=golang
 *
 * [16] 最接近的三数之和
 */

// @lc code=start
func threeSumClosest(nums []int, target int) int {
	var allSum []int
	len := len(nums)
	for i, Enum := range nums {
		for i2 := i + 1; i2 < len; i2++ {
			for i3 := i2 + 1; i3 < len; i3++ {
				allSum = append(allSum, Enum+ nums[i2] + nums[i3])
			}
		}
	}
	ans := 1000
	closest := 1000

	for _, num := range allSum {
		if math.Abs(float64(target - num)) < float64(closest) {
			ans = num
			closest = int(math.Abs(float64(target - num)))
		}
	}

	return ans
}

func main() {
	println(threeSumClosest([]int{
	-11,-2,17,-16,1,-5,-5,-5,-20,7,10,-2,3,-7,-17,-13,-19,-15,-8,-7,6,-6,-8,-4,12,-12,9,-17,-13,4,-5,-15,-9,-18,-17,1,-15,-8,14,8,20,-3,-11,17,-18,10,-16,5,-9,-18,2,-3,4,-18,2,20,0,-6,18,-12,0,-17,3,-19,-20,15,12,-17,-7,8,16,7,-5,5,-13,16,-18,-7,-9,-8,-17,6,-18,0,-15,10,-13,7,9,20,7,-13,3,0,0,19,8,0,-5,-9,6,8,16,14,3,-4,5,9,-12,-19,16,6,

	}, -48))
}
// @lc code=end

