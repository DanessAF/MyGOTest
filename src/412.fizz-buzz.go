package main

import "strconv"

/*
 * @lc app=leetcode.cn id=412 lang=golang
 *
 * [412] Fizz Buzz
 */

// @lc code=start
func fizzBuzz(n int) []string {
	var result []string
	for i := 1; i <= n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i))
		}
		println()
	}
	return result
}
// func main() {
// 	fizzBuzz(11)

// }
// @lc code=end

