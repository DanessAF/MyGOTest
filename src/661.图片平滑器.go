package main

import "fmt"

/*
 * @lc app=leetcode.cn id=661 lang=golang
 *
 * [661] 图片平滑器
 */

// @lc code=start
func imageSmoother(img [][]int) [][]int {
	xLen := len(img) - 1
	yLen := len(img[0]) - 1
	var ans [][]int
	for _, _ = range img {
		ans = append(ans, make([]int, yLen + 1))
	}


	for x := 0; x <= xLen; x++ {
		for y := 0; y <= yLen; y++ {
			use := 0
			tmpXRow := []int{x - 1, x, x + 1}
			tmpYRow := []int{y - 1, y, y + 1}
			smoother := 0
			for _, xPosition := range tmpXRow {
				for _, yPosition := range tmpYRow {
					if (0 <= xPosition && xPosition <= xLen) && (0 <= yPosition && yPosition <= yLen) {
						use++
						smoother += img[xPosition][yPosition]
					}
				}
			}
			ans[x][y] = smoother / use
		}
	}
	return ans
}

func main() {
	t := imageSmoother([][]int{
		{100,200,100},
		{200,50,200},
		{100,200,100},
	})
	fmt.Printf("%v", t)
}
// @lc code=end

