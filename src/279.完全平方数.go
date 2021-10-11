package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

// @lc code=start
func numSquares(n int) int {
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j * j <= i; j++ {
			if minn > f[i-j*j] {
				minn = f[i - j * j]
			}
		}
		println(minn)
		f[i] = minn + 1
		fmt.Printf("%v", f)
		println()
	}
	return f[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	println(numSquares(3))
}
// @lc code=end

