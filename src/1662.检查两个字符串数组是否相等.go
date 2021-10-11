package main

import "strings"

/*
 * @lc app=leetcode.cn id=1662 lang=golang
 *
 * [1662] 检查两个字符串数组是否相等
 */

// @lc code=start
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	left := ""
	right := ""

	leftLen := len(word1)
	rightLen := len(word2)

	leftPosition := 0
	rightPosition := 0

	for leftLen > leftPosition || rightLen > rightPosition {
		if rightLen > rightPosition {
			rightPosition++
			right += word2[rightPosition - 1]
		}
		if leftLen > leftPosition {
			leftPosition++
			left += word1[leftPosition - 1]
		}
		leftStrLen := strings.Count(left, "") - 1
		rightStrLen := strings.Count(right, "") - 1

		if leftStrLen == rightStrLen && left != right {
			return false
		}
	}

	if left != right {
		return false
	}
	return true
}

func main() {
	println(arrayStringsAreEqual([]string{"abc","d","defg"}, []string{"abcddef"}))
}
// @lc code=end

