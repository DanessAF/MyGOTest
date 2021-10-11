package main

import "fmt"

/*
 * @lc app=leetcode.cn id=216 lang=golang
 *
 * [216] 组合总和 III
 */




// @lc code=start
func combinationSum3(k int, n int) (combinationSum[][]int) {
	var backtracking func(k int, n int, idx int, track []int, sum int)
	backtracking = func(k int, n int, idx int, track []int, sum int) {
		if len(track) > k {
			return
		}

		for i:= idx; i <= 9; i++{
			track = append(track,i)
			sum += i
			if len(track) == k && sum == n  {
				temp := make([]int,k)
				copy(temp,track)
				combinationSum = append(combinationSum,temp)
				continue
			}
			if sum < n {
				backtracking(k, n, i + 1, track, sum)
			}
			track=track[:len(track)-1]
			sum -= i
		}
	}

	backtracking(k, n, 1, []int{}, 0)
	return
}

func main() {
	fmt.Printf("%v", combinationSum3(3, 7))
}
// @lc code=end