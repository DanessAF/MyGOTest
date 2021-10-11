package main

import "strings"

/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	if strsLen == 1 {
		return strs[0]
	}

	longestCommonPrefix := strs[0]
	longestCommonPrefixLen := strings.Count(longestCommonPrefix, "") - 1
	for i := 0; i <= longestCommonPrefixLen; i++ {
		for strNow := 0; strNow < strsLen; strNow++ {
			strNowLen := strings.Count(strs[strNow], "") - 1
			if i > strNowLen {
				return longestCommonPrefix[:i - 1]
			}
			if longestCommonPrefix[:i] != strs[strNow][:i] {
				return longestCommonPrefix[:i - 1]
			}
		}

	}
	return longestCommonPrefix
}

// func main() {
// 	a := []string{"flower","flower","flower","flower"}
// 	println(longestCommonPrefix(a))
// }

// @lc code=end

