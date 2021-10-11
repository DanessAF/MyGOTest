package main
/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	var ans [][]int
	ans = append(ans, []int{1})

	for i := 1; i <= numRows; i++ {
		ans = append(ans, []int{})
		for y := 0; y < i; y++ {
			left := 0
			right := 0
			if y == 0 {
				left = 0
			} else {
				left = ans[i - 1][y - 1]
			}
			if y < len(ans[i - 1]) {
				right = ans[i - 1][y]
			} else {
				right = 0
			}

			ans[i] = append(ans[i], left + right)
		}
	}
	return ans[1:]

}

// func main() {
// 	result := generate(5)
// 	println(len(result))
// 	for _, i2 := range result {
// 		str := ""
// 		for _, i3 := range i2 {
// 			println(i3)
// 		}
// 		println(str)
// 	}
// }
// @lc code=end

