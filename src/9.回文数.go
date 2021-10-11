package main

import (
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	a := strconv.Itoa(x)
	aLen := strings.Count(a, "") - 2
	for i := 0; i <= aLen / 2; i++ {
		println(a[i], a[aLen - i])
		if a[i] != a[aLen - i] {
			return false
		}
	}
	return true
}

// @lc code=end

