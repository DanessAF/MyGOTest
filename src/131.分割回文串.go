package main

import (
	"fmt"
	"strings"
)

func partition(s string) (res [][]string) {
	var partition []string
	sLen := strings.Count(s, "") - 1
	for _, str := range s {
		partition = append(partition, string(str))
	}

	res = append(res, partition)

	if sLen == 2 {
		if confirm(s) {
			res = append(res, []string{partition[0] + partition[1]})
		}
		return
	}
	var dfs = func(idx int, sLen int, partition []string) {}
	dfs = func(left int, sLen int, partition []string) {
		if left > sLen {
			return
		}


		temp := make([]string, len(partition))
		copy(temp, partition)
		partitionLen := len(partition)

		for i := left + 1; i < partitionLen; i++ {
			dfs(i, len(temp), temp)

			temp[left] += string(partition[i])
			temp = append(temp[:left + 1], partition[i + 1:]...)
			fmt.Printf("%v", temp)
			println()
			if !confirm(temp[left]) {
				continue
			}

			for i2 := 0; i2 < left; i2++ {
				if !confirm(temp[i2]) {
					continue
				}
			}

			re := make([]string, len(temp))
			copy(re, temp)
			res = append(res, re)
			dfs(i, len(re), re)
		}
	}

	dfs(0, sLen, partition)
	return
}

func confirm(str string) bool {
	strLen := strings.Count(str, "") - 1

	for i := 0; i <= strLen / 2; i++ {
		if str[i] != str[strLen - i - 1] {
			return false
		}
	}
	return true
}


func main()  {
	fmt.Printf("%v", partition("cbbbcc"))
}