package main

import (
	"fmt"
	"sort"
	"strconv"
)

/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
func fourSum(nums []int, target int) [][]int {
	fined := make(map[string]string)
	sort.Ints(nums)

	var ans [][]int
	len := len(nums)

	//ed := 0
	for i, left := range nums {
		for i2 := i + 1; len > i2 ; i2++ {
			for i3 := i2 + 1; len > i3 ; i3++ {
				for i4 := i3 + 1; len > i4 ; i4++ {
					if left + nums[i2] + nums[i3] + nums[i4] == target {
						//if nums[i4] == ed {
						//	continue
						//}
						tmp := []int{left, nums[i2], nums[i3], nums[i4]}
						findKey := strconv.Itoa(left)  + "|" + strconv.Itoa(tmp[1]) + "|" + strconv.Itoa(tmp[2]) + "|" + strconv.Itoa(tmp[3])
						_, exist := fined[findKey]
						if exist {
							break
						}
						fined[findKey] = "1"
						ans = append(ans, tmp)
						//ed = nums[i4]
					}
				}
			}
			i++
		}
	}

	return ans
}

func main() {
	fmt.Printf("%v", fourSum([]int{
	1,0,-1,0,-2,2,

	//2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,

	}, 0))
}
// @lc code=end

