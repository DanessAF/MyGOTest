package main

import "strings"

/*
 * @lc app=leetcode.cn id=5 lang=golang
 *
 * [5] 最长回文子串
 */

// @lc code=start
func longestPalindrome(s string) string {
	var left string
	var right string
	var bright string

	var Palindrome []string
	sLen := strings.Count(s, "") - 1
	if sLen == 1 {
		return s
	}
	longestPalindrome := 0

	for i := 0; i < sLen; i++ {
		left = ""
		right = ""
		bright = ""

		for i2 := i; i2 >= 0; i2-- {
			if sLen < i2 + i {
				continue
			}
			left = s[i- i2:i]
			right = s[i:i+ i2]

			if i+ i2 + 1 <= sLen {
				bright = s[i:i+ i2 + 1]
			}

			if  confirms(left, bright) {
				if strings.Count(left + bright, "") - 1 > longestPalindrome  {
					longestPalindrome = strings.Count(left + bright, "") - 1
				}
				Palindrome = append(Palindrome, left + bright)
				break
			} else if confirm(left, right) {
				if strings.Count(left + right, "") - 1 > longestPalindrome  {
					longestPalindrome = strings.Count(left + right, "") - 1
				}
				Palindrome = append(Palindrome, left + right)
				break
			}
		}
	}

	result := ""
	for _, s2 := range Palindrome {
		if longestPalindrome == strings.Count(s2, "") - 1 {
			result = string(s2)
		}
	}

	if result == "" {
		return string(s[0])
	}
	return result
}

func confirmLength(left string, bright string, longestPalindrome int) bool {
	if bright == "" {
		return false
	}
	if strings.Count(left + bright, "") - 1 < longestPalindrome {
		return true
	}

	return false
}

func confirm(left string, right string) bool {
	leftLen := strings.Count(left, "") - 2
	result := true

	for i := 0; i <= leftLen; i++ {
		if left[i] != right[leftLen - i] {
			return false
		}
	}
	return result
}

func confirms(left string, bright string) bool {
	if strings.Count(bright, "") - 1 < 2 {
		return false
	}
	return confirm(left, bright[1:])
}

func main() {
	println(longestPalindrome("fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffgggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggggg"))
}
// @lc code=end

