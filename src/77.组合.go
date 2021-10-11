package main

import "fmt"

var res [][]int
/*
 * @lc app=leetcode.cn id=77 lang=golang
 *
 * [77] 组合
 */

// @lc code=start
func combine(n int, k int) [][]int {

	backtracking(k, n, 1, []int{}, 0)
	return res
}

func backtracking(n int, k int, idx int, track []int, sum int)  {
	//fmt.Printf("%v", track)

	if len(track) > k {
		return
	}

	for i:= idx; i <= 9; i++{
		track = append(track,i)
		sum += i
		if len(track) == k && sum == n  {
			temp := make([]int,k)
			copy(temp,track)
			res = append(res,temp)
			continue
		}
		if sum < n {
			backtracking(n, k, i + 1, track, sum)
		}
		track=track[:len(track)-1]
		sum -= i
	}
}

func main() {
	fmt.Printf("%v", combine(3, 9))
}
// @lc code=end

