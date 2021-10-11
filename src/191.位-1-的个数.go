package main
/*
 * @lc app=leetcode.cn id=191 lang=golang
 *
 * [191] 位1的个数
 */

// @lc code=start
func hammingWeight(num uint32) int {
	len := 0
	i := 1
	for num != 0 {
		len += int(num) & i
		num = num >> 1
	}
	return  len
}

func main()  {
	println(hammingWeight(00000000000000000000000000001011))
}
// @lc code=end

