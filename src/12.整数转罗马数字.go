package main

/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */

// @lc code=start
func intToRoman(num int) string {
	mm := map[string]int{
		"M": 1000,
		"CM": 900,
		"D": 500,
		"CD": 400,
		"C": 100,
		"XC": 90,
		"L": 50,
		"XL": 40,
		"X": 10,
		"IX": 9,
		"V": 5,
		"IV": 4,
		"I": 1,
	}
	var mv []string
	mv = append(mv, "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I")
	var ans []string

	for _, m := range mv{
		println(m)
		for num >= mm[m] {
			num -= mm[m]
			println(num)
			ans = append(ans, m)
		}
		if num == 0 {
			 break
		}
	}
	a := ""
	for _, an := range ans {
		a += an
	}

	return a
}

func main() {
	println(intToRoman(3999))
}
// @lc code=end

