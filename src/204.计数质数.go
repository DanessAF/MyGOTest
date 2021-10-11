package main

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	var Primes []int
	for i := 1; i < n; i++ {
		Primes = append(Primes, i)
	}

	if len(Primes) >= 1 {
		Primes[0] = 0
	}

	for _, prime := range Primes {
		if prime == 1 || prime == 0 {
			continue
		}
		for i2 := 2; i2 < n - 1; i2++ {
			noPrime := prime * i2
			if noPrime <= n - 1 {
				Primes[noPrime - 1] = 0
			} else {
				i2 = n
			}
		}
	}

	result := 0
	for _, prime := range Primes {
		if prime != 0 {
			println(prime)
			result++
		}
	}
	return result
}

func main() {
	println(countPrimes(100))
}
// @lc code=end

