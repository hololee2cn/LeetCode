package main

import "fmt"

func minSubarray(nums []int, p int) int {
	n := len(nums)
	sum := 0
	m := make(map[int]int)
	m[0] = -1
	for _, v := range nums {
		sum += v
	}
	x := sum % p
	if x == 0 {
		return 0
	}
	res := n
	sum = 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		y := sum % p
		// s1代表数组总和,s2代表符合条件的子数组总和 s1 mod p == s2 mod p == x => (s1-s2) mod p == 0
		// s2又能用前缀和表示,i代表子数组的右边界,j代表子数组的左边界,有f[i]-f[j] = s2，进而 (f[i]-f[j]) mod p == x
		// f[j] mod p == (f[i]-x) mod p 0 <j <= i，我们可以用hash去记录i位置之前的f[j]前缀和 %p的记录
		// 如果存在则表示有符合条件的子数组,子数组范围为i-j
		// i 之前的 f[j]前缀和 f[i] mod p = (f[j]-x) mod p 如果存在则更新 i-j+1子数组大小
		k := (y - x + p) % p
		if _, ok := m[k]; ok {
			res = min(res, i-m[k])
		}
		m[y] = i
	}

	if res == n {
		return -1
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type minSubarrayInput struct {
	nums []int
	p    int
}

func main() {
	cases := []minSubarrayInput{
		{
			nums: []int{3, 1, 4, 2},
			p:    6,
		},
		{
			nums: []int{6, 3, 5, 2},
			p:    9,
		},
		{
			nums: []int{1, 2, 3},
			p:    3,
		},
	}
	for _, v := range cases {
		fmt.Println(minSubarray(v.nums, v.p))
	}
}
