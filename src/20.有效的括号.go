package main

/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	tmp := ""
	valid := make(map[string]string)
	valid["{"] = "}" // 赋值
	valid["("] = ")" // 赋值
	valid["["] = "]"
	result := false

	for i := 0; i < len(s); i++ {
		str := string(s[i])

		if valid[str] != "" {
			tmp += str
			continue
		}

		if tmp == "" {
			return false
		}
		tmpLen := len(tmp)

		if tmpLen > len(s[i:]) {
			return false
		}

		if string(s[i]) == valid[string(tmp[tmpLen - 1])] {
			tmp = tmp[:tmpLen - 1]
			result = true
		} else {
			return false
		}
	}
	if tmp != "" {
		return false
	}
	println(1)
	return result
}

// func main() {
// 	println(isValid("([]){"))
// }
// @lc code=end

