package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) (res []string) {
	generateParentheses := func(idx int, now string, n int){}
	generateParentheses = func(idx int, now string, n int) {
		if n * 2 == idx {
			if isParenthesis(now) {
				res = append(res, now)
			}
			return
		}
		if  {

		}
		generateParentheses(idx + 1, now + "(", n)
		generateParentheses(idx + 1, now + ")", n)
	}
	generateParentheses(1, "(", n)
	return
}

func isParenthesis(str string) bool {
	now := ""
	for _, s := range str {
		if s == '(' {
			now += "()"
		}
		if s == ')' {
			if len(now) == 0 {
				return false
			}
			now = now[:len(now) - 2]
		}
	}
	if len(now) > 0 {
		return false
	}
	return true
}


func main() {
	fmt.Printf("%v", generateParenthesis(3))
}
// @lc code=end

