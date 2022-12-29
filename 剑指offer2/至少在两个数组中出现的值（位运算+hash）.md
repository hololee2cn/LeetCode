# 至少在两个数组中出现的值（位运算+hash）

```go
package main

import "fmt"

// 给你三个整数数组 nums1、nums2 和 nums3 ，请你构造并返回一个 元素各不相同的 数组，且由 至少 在 两个 数组中出现的所有值组成。数组中的元素可以按 任意 顺序排列。
//
// 示例 1：
//
// 输入：nums1 = [1,1,3,2], nums2 = [2,3], nums3 = [3]
// 输出：[3,2]
// 解释：至少在两个数组中出现的所有值为：
// - 3 ，在全部三个数组中都出现过。
// - 2 ，在数组 nums1 和 nums2 中出现过。
//
// 示例 2：
//
// 输入：nums1 = [3,1], nums2 = [2,3], nums3 = [1,2]
// 输出：[2,3,1]
// 解释：至少在两个数组中出现的所有值为：
// - 2 ，在数组 nums2 和 nums3 中出现过。
// - 3 ，在数组 nums1 和 nums2 中出现过。
// - 1 ，在数组 nums1 和 nums3 中出现过。
//
// 示例 3：
//
// 输入：nums1 = [1,2,2], nums2 = [4,3,3], nums3 = [5]
// 输出：[]
// 解释：不存在至少在两个数组中出现的值。
//
// 提示：v & (v - 1)解释 1.三个数组出现：111&110=110(√) 2.两个数组出现：110&101=100(√),101&100=100(√),011&010=010(√) 3.一个数组出现：100&011=000(×),010&001=000(×),001&000=000(×),

func twoOutOfThree(nums1, nums2, nums3 []int) (ans []int) {
	mask := map[int]int{}
	for i, nums := range [][]int{nums1, nums2, nums3} {
		for _, x := range nums {
			mask[x] |= 1 << i
		}
	}
	for x, m := range mask {
		if m&(m-1) > 0 {
			ans = append(ans, x)
		}
	}
	return
}

type twoOutOfThreeInput struct {
	nums1 []int
	nums2 []int
	nums3 []int
}

func main() {
	cases := []twoOutOfThreeInput{
		{
			[]int{1, 1, 3, 2},
			[]int{2, 3},
			[]int{3},
		},
		{
			[]int{3, 1},
			[]int{2, 3},
			[]int{1, 2},
		},
		{
			[]int{1, 2, 2},
			[]int{4, 3, 3},
			[]int{5},
		},
	}
	for _, v := range cases {
		fmt.Println(twoOutOfThree(v.nums1, v.nums2, v.nums3))
	}
}
```