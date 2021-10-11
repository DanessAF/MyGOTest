package main
/*
 * @lc app=leetcode.cn id=190 lang=golang
 *
 * [190] 颠倒二进制位
 */

// @lc code=start
func reverseBits(num uint32) uint32 {
	var ans uint32 = 0
	for i := 0; i < 31; i++ {
		ans <<= 1
		ans += num & 1
		num >>= 1
	}
	return  ans
}

func main() {
	var num uint32 =   3221225470
	println(reverseBits(num))
}
// @lc code=end

