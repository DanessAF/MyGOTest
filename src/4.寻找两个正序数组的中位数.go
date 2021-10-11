package main

import (
	"math/rand"
	"time"
)

func main(){
	nums3 := ArrayFill(999)
	println(nums3[len(nums3)/ 2 - 1])
	println(nums3[len(nums3)/ 2])
	nums1, nums2 := ArrayDist(nums3)

	println(len(nums1), len(nums2))
	println(findMedianSortedArrays(nums1, nums2))
}

/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	numsCount1 := len(nums1)
	numsCount2 := len(nums2)
	right := numsCount1 + numsCount2
	count := numsCount1 + numsCount2

	removed := 0
	if nums1[len(nums1) / 2] < nums2[len(nums2) / 2] {
		mid := findFirstLT(nums2[len(nums2) / 2], nums1)
		removed = len(nums1[mid - 1:]) + len(nums2[:mid])
		if removed > count / 2 {
			mid = findFirstLT(nums1[len(nums1) / 2], nums2)
			removed =  len(nums1[:mid]) + len(nums2[mid - 1:])

			nums1 = nums1[:(len(nums1) / 2) + 1]
			nums2 = nums2[:mid - 1]
		} else {
			nums1 = nums1[:mid - 1]
			nums2 = nums2[:(len(nums2) / 2) + 1]
		}
	} else {
		mid := findFirstLT(nums1[len(nums1) / 2], nums2)
		removed =  len(nums1[:len(nums1) / 2]) + len(nums2[mid - 1:])
		if removed > count / 2 {
			mid = findFirstLT(nums2[len(nums2) / 2], nums1)

			removed =  len(nums1[mid - 1:]) + len(nums2[:mid])
			nums1 = nums1[:mid - 1]
			nums2 = nums2[:(len(nums2) / 2) + 1]
		} else {
			nums1 = nums1[:(len(nums1) / 2) + 1]
			nums2 = nums2[:mid - 1]
		}
	}

	println(removed)
	println(right - count / 2)
	println(len(nums1))
	println(len(nums2))

	return float64(1)
}


func findFirstLT(find int, nums []int) int {
	start := 0
	end := len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] <= find {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return start
}




func binarySearch(find int, nums []int) int {
	start := 0
	end := len(nums)
	for start < end {
		mid := (start + end) / 2
		if nums[mid] == find {
			return mid
		}
		if nums[mid] > find {
			start = mid + 1
		} else {
			end = mid
		}
	}
	return -1
}

func ArrayDist(nums []int) ([]int, []int) {
	var nums1  []int
	var nums2  []int

	rand.Seed(time.Now().Unix())
	for _, num := range nums {
		if rand.Intn(100000) % 20 == 1 {
			nums1 = append(nums1, num)
		} else {
			nums2 = append(nums2, num)
		}
	}
	return nums1,nums2
}

func ArrayFill(num int) []int {
	var m []int
	var i int
	for i = 0; i <= num; i++ {
		if len(m) < 1 {
			m = append(m,rand.Intn(100))
		} else {
			m = append(m, m[i -1] + rand.Intn(100))
		}
	}
	return m
}
// @lc code=end

