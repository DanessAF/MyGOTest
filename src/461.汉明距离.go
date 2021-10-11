
package main

/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) int {
	len := 0
	i := 1
	for x != 0 || y != 0 {
		if !(x & i == y & i) {
			len++
		}
		x = x >> 1
		y = y >> 1
	}
	return  len
}

func main()  {
	println(hammingDistance(1, 4))
}
// @lc code=end

