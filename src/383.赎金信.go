package main

/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	countryCapitalMap := make(map[string]int)

	for _, mag := range magazine {
		_, exist := countryCapitalMap[string(mag)]

		if exist {
			countryCapitalMap[string(mag)]++
		} else {
			countryCapitalMap[string(mag)] = 1
		}
	}

	for _, note := range ransomNote {
		if countryCapitalMap[string(note)] == 0 {
			return false
		}
		countryCapitalMap[string(note)]--
	}
	return true
}

func main() {
	println(canConstruct("a", "b"))
}
// @lc code=end

