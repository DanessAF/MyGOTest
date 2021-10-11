package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(lengthOfLongestSubstring("\"abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!\\\\\"#$%&\\'()*+,-./:;<=>?@[\\\\\\\\]^_`{|}~ "))
}
/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	est := make(map[string]int)

	result := 0
	slen := strings.Count(s, "") - 1
	for i := 0; i < slen; i++ {
		nowString := string(s[i])
		if nowString == " " {
			nowString = "ab"
		}

		if est[nowString] == 0 {
			est[nowString] = i + 1
			if  len(est) > result {
				result = len(est)
			}
		} else {
			if  len(est) > result {
				result = len(est)
			}
			i = est[nowString] - 1
			est = make(map[string]int)
		}
	}
	return result
}
// @lc code=end

