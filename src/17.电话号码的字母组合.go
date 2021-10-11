package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */

// @lc code=start
func letterCombinations(digits string) []string {
	making := make(map[string][]string)
	making["2"] = append(making["2"], []string{"a", "b", "c"}...)
	making["3"] = append(making["3"], []string{"d", "e", "f"}...)
	making["4"] = append(making["4"], []string{"g", "h", "i"}...)
	making["5"] = append(making["5"], []string{"j", "k", "l"}...)
	making["6"] = append(making["6"], []string{"m", "n", "o"}...)
	making["7"] = append(making["7"], []string{"p", "q", "r", "s"}...)
	making["8"] = append(making["8"], []string{"t", "u", "v"}...)
	making["9"] = append(making["9"], []string{"w", "x", "y", "z"}...)

	var ans []string
	if digits == "" {
		return ans
	}

	ans = append(ans,  making[string(digits[0])]...)
	digitsLen := strings.Count(digits, "") - 1

	for i := 1; i < digitsLen; i++ {
		needMaking := making[string(digits[i])]
		var tmp []string
		for _, an := range ans {
			for _, s := range needMaking {
				tmp = append(tmp, an + s)
			}
		}
		ans = tmp
	}
	return ans
}

func main() {
	fmt.Printf("%v", letterCombinations("23"))

}
// @lc code=end

