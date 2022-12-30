# 最长连续序列（并查集/哈希）

## 哈希 version
```go
package main

import "fmt"

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// 
// 请你设计并实现时间复杂度为O(n) 的算法解决此问题。
// 示例 1：
// 
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
// 示例 2：
// 
// 输入：nums = [0,3,7,2,5,8,4,6,0,1]
// 输出：9

func longestConsecutive(nums []int) int {
	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}
	longestRes := 0
	for num := range numSet {
		if !numSet[num-1] {
			cur := num
			curLen := 1
			for numSet[cur+1] {
				cur++
				curLen++
			}
			if longestRes < curLen {
				longestRes = curLen
			}
		}
	}
	return longestRes
}

func main() {
	cases := [][]int{
		{
			100, 4, 200, 1, 3, 2,
		},
		{
			0, 3, 7, 2, 5, 8, 4, 6, 0, 1,
		},
	}
	for _, v := range cases {
		fmt.Println(longestConsecutive(v))
	}
}
```

## 并查集 version
```go
package main

import "fmt"

func longestConsecutive2(nums []int) int {
	un := newUn(len(nums))
	m := make(map[int]int)
	for i := range nums {
		if _, ok := m[nums[i]]; ok { // 存在重复元素
			continue
		}
		if _, ok := m[nums[i]-1]; ok {
			un.union(i, m[nums[i]-1])
		}
		if _, ok := m[nums[i]+1]; ok {
			un.union(i, m[nums[i]+1])
		}
		m[nums[i]] = i
		fmt.Printf("nums[%d]:%d,parent:%+v,rank:%+v\n", i, nums[i], un.parent, un.rank)
	}
	fmt.Printf("src:%d,dst:%d,isConnect:%v\n", 1, 4, un.connect(1, 4))
	return un.getMaxConnection()
}

type un struct {
	parent []int
	rank   []int
}

func newUn(n int) *un {
	parent := make([]int, n)
	size := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &un{
		parent: parent,
		rank:   size,
	}
}

func (u *un) find(n int) int {
	if u.parent[n] != n {
		u.parent[n] = u.find(u.parent[n])
	}
	return u.parent[n]
}

func (u *un) union(x, y int) {
	rootx := u.find(x)
	rooty := u.find(y)
	if rootx != rooty {
		u.parent[rootx] = rooty
		u.rank[rooty] += u.rank[rootx]
	}
}

func (u *un) connect(src int, dst int) bool {
	return u.find(src) == u.find(dst)
}

func (u *un) getMaxConnection() int {
	maxRes := 0
	maxInt := func(v1, v2 int) int {
		if v1 < v2 {
			return v2
		}
		return v1
	}
	for i := range u.parent {
		if u.parent[i] == i {
			maxRes = maxInt(maxRes, u.rank[i])
		}
	}
	return maxRes
}

func main() {
	cases := [][]int{
		{100, 4, 200, 1, 3, 2},
		{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
	}
	for _, v := range cases {
		fmt.Println(longestConsecutive2(v))
	}
}
```
