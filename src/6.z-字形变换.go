package main

import (
	"strings"
)

/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows == 1{
		return s
	}
	zRow := make([][]string, numRows)
	for i := range zRow {
		zRow[i] = make([]string, strings.Count(s, ""))
	}

	x := 0
	y := 0
	numRos := numRows - 1
	for _, str := range s {
		println(x)
		println(y)
		zRow[x][y] = string(str)
		if x == numRos {
			y++
			x--
			continue
		} else if y % numRos == 0 {
			x++
			continue
		} else {
			y++
			x--
		}
	}

	ans := ""
	for _, yRow := range zRow {
		for _, str := range yRow {
			if str != "" {
				ans += str
			}
		}
	}

	return ans
}

func main()  {
	println(convert("AB", 1))
}
// @lc code=end

