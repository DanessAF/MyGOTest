package main


/*
 * @lc app=leetcode.cn id=13 lang=golang
 *
 * [13] 罗马数字转整数
 */

// @lc code=start
func romanToInt(s string) int {
	m := map[string]int{"I": 1,
		"V": 5, "X": 10,
		"L": 50, "C": 100, "D": 500, "M": 1000,
	}
	var c []int
	var result = 0

	for _, str := range s {
		c = append(c, m[string(str)])
	}

	clen := len(c)
	for i := clen - 1; i >= 0; i-- {
		int := c[i]

		if i == 0 {
			result += int
			continue
		}

		if int > c[i - 1] {
			result += int
			result -= (c[i - 1] * 2)
		} else {
			result += int
		}
	}
	return result
}

func main() {
	end := romanToInt("LVIII")
	println()
	println(end)
}
// @lc code=end

