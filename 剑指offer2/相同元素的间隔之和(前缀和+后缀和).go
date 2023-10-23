package main

import "fmt"

// 给你一个下标从 0 开始、由 n 个整数组成的数组 arr 。
//
// arr 中两个元素的 间隔 定义为它们下标之间的 绝对差 。更正式地，arr[i] 和 arr[j] 之间的间隔是 |i - j| 。
//
// 返回一个长度为 n 的数组 intervals ，其中 intervals[i] 是 arr[i] 和 arr 中每个相同元素（与 arr[i] 的值相同）的 间隔之和 。
//
// 注意：|x| 是 x 的绝对值。
//
// 示例 1：
//
// 输入：arr = [2,1,3,1,2,3,3]
// 输出：[4,2,7,2,4,4,5]
// 解释：
// - 下标 0 ：另一个 2 在下标 4 ，|0 - 4| = 4
// - 下标 1 ：另一个 1 在下标 3 ，|1 - 3| = 2
// - 下标 2 ：另两个 3 在下标 5 和 6 ，|2 - 5| + |2 - 6| = 7
// - 下标 3 ：另一个 1 在下标 1 ，|3 - 1| = 2
// - 下标 4 ：另一个 2 在下标 0 ，|4 - 0| = 4
// - 下标 5 ：另两个 3 在下标 2 和 6 ，|5 - 2| + |5 - 6| = 4
// - 下标 6 ：另两个 3 在下标 2 和 5 ，|6 - 2| + |6 - 5| = 5
// 示例 2：
//
// 输入：arr = [10,5,10,10]
// 输出：[5,0,3,4]
// 解释：
// - 下标 0 ：另两个 10 在下标 2 和 3 ，|0 - 2| + |0 - 3| = 5
// - 下标 1 ：只有这一个 5 在数组中，所以到相同元素的间隔之和是 0
// - 下标 2 ：另两个 10 在下标 0 和 3 ，|2 - 0| + |2 - 3| = 3
// - 下标 3 ：另两个 10 在下标 0 和 2 ，|3 - 0| + |3 - 2| = 4

func getDistances(arr []int) []int64 {
	n := len(arr)
	// 前缀和转移方程 r[i] = r[j] + (d[i]-d[j])*个数 表示i位置前缀所有出现s[i]数字在j位置出现s[i]数字个数+（i-j）距离 * 出现s[i]数字的个数
	m := make(map[int]*[2]int)
	res := make([]int, n)
	// init
	// 前缀
	for i := 0; i < n; i++ {
		if _, ok := m[arr[i]]; !ok {
			m[arr[i]] = &[2]int{i, 0}
		}
		// 统计前面一个出现同样数字位置的个数
		if m[arr[i]][1] != 0 {
			res[i] = res[m[arr[i]][0]] + getDistancesAbs(i-m[arr[i]][0])*m[arr[i]][1]
		}
		// 标注上一个位置
		m[arr[i]][0] = i
		// 标注上一个位置出现的个数
		m[arr[i]][1]++
	}
	// 后缀
	res2 := make([]int, n)
	m2 := make(map[int]*[2]int)
	for i := n - 1; i >= 0; i-- {
		if _, ok := m2[arr[i]]; !ok {
			m2[arr[i]] = &[2]int{i, 0}
		}
		// 统计后面一个出现同样数字位置的个数
		if m[arr[i]][1] != 0 {
			res2[i] = res2[m2[arr[i]][0]] + getDistancesAbs(i-m2[arr[i]][0])*m2[arr[i]][1]
		}
		m2[arr[i]][0] = i
		m2[arr[i]][1]++
	}
	// 最后将前缀和后缀进行相加
	ret := make([]int64, n)
	for i := 0; i < n; i++ {
		ret[i] += int64(res[i] + res2[i])
	}
	return ret
}

func getDistancesAbs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func main() {
	cases := [][]int{
		{2, 1, 3, 1, 2, 3, 3},
		{10, 5, 10, 10},
	}
	for _, v := range cases {
		fmt.Println(getDistances(v))
	}
}
