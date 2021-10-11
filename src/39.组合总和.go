package main

import "fmt"

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */


// @lc code=start
func combinationSum(candidates []int, target int) (ans [][]int) {
	comb := []int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {

		println(target, idx)
		fmt.Printf("%v", comb)
		println()
		//println()
		//fmt.Printf("%v", ans)
		if idx == len(candidates) {
			return
		}
		if target == 0 {
			ans = append(ans, append([]int(nil), comb...))
			return
		}
		// 直接跳过
		dfs(target, idx+1)
		// 选择当前数
		if target-candidates[idx] >= 0 {
			comb = append(comb, candidates[idx])
			dfs(target-candidates[idx], idx)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return
}



func main() {
	fmt.Printf("%v", combinationSum([]int{2,3,6,7}, 7))
}
// @lc code=end

