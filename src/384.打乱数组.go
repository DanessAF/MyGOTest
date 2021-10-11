package main
import "math/rand"

/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
type Solution struct {
	data []int
}

func Constructor(nums []int) Solution {
	solution := Solution{
		data: nums,
	}
	return solution
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.data
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.data))
	copy(nums, this.data)
	rand.Shuffle(len(nums), func(i int, j int) {
		println(i)
		nums[i], nums[j] = nums[j], nums[i]
	})
	return nums
}


func main() {
	t := Constructor([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	t.Shuffle()
}


/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

