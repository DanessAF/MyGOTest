package main

/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
func findLongestWord(s string, dictionary []string) string {
	var finedWord []string
	var max int

	for _, s2 := range dictionary {
		ps := 0
		ls := len(s2)
		for _, i2 := range s {
			if string(i2) == string(s2[ps]) {
				ps++
			}
			if ps == ls {
				finedWord = append(finedWord, s2)
				if max < len(s2) {
					max = len(s2)
				}
				break
			}

		}
	}

	var maxFind[]string
	for _, s2 := range finedWord {
		if len(s2) == max {
			maxFind = append(maxFind, s2)
		}
	}

	min := ""
	for _, s2 := range maxFind {
		if min == ""{
			min = s2
		}
		if s2 < min {
			min = s2
		}
	}

	return min
}

func main() {
	println(findLongestWord("abce", []string{"abc", "abe"}))
}

// @lc code=end
