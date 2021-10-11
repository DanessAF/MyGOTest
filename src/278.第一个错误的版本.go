package main

/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 *
 */

func isBadVersion(version int) bool {
	return version >= 6
}
func firstBadVersion(n int) int {
	start := 0
	end := n
	for start < end {
		mid := (start + end) / 2
		println(start)
		println(end)
		println(mid)
		if !isBadVersion(mid) {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}

func main() {
	print(firstBadVersion(10))
}
// @lc code=end

