package main

/*
 * @lc app=leetcode.cn id=1252 lang=golang
 *
 * [1252] 奇数值单元格的数目
 */

// @lc code=start
func oddCells(m int, n int, indices [][]int) int {
	var ans [][]int
	oddCells := 0
	for mm := 0; mm < m; mm++ {
		ans = append(ans, []int{})
		for nn := 0; nn < n; nn++ {
			ans[mm] = append(ans[mm], 0)
		}
	}

	for _, index := range indices {
		ri := index[0]
		ci := index[1]
		println(ri, ci)
		for mm := 0; mm < m; mm++ {
			ans[mm][ci]++
		}
		for nn := 0; nn < n; nn++ {
			ans[ri][nn]++
		}
	}

	for mm := 0; mm < m; mm++ {
		ans = append(ans, []int{})
		for nn := 0; nn < n; nn++ {
			if ans[mm][nn]%2 == 1 {
				oddCells++
			}
		}
	}

	return oddCells
}

func main() {
	println(oddCells(2, 2, [][]int{[]int{1,1}, []int{0, 0}}))
}
// @lc code=end

