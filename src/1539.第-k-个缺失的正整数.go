package main
/*
 * @lc app=leetcode.cn id=1539 lang=golang
 *
 * [1539] 第 k 个缺失的正整数
 */

// @lc code=start
func findKthPositive(arr []int, k int) int {
	count := len(arr) + k
	lack := 0
	for i := 1; i <= count; i++ {
		if len(arr) == 0 {
			lack++
		} else if arr[0] != i {
			lack++
		} else {
			arr = arr[1:]
		}
		if lack == k {
			return i
		}

	}
	return 1
}

// func main() {
// 	i := []int{1, 2, 3, 4}
// 	println(findKthPositive(i, 2))

// }
// @lc code=end

