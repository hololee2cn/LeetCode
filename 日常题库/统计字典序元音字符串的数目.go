package main

import "fmt"

// 给你一个整数 n，请返回长度为 n 、仅由元音 (a, e, i, o, u) 组成且按 字典序排列 的字符串数量。
//
// 字符串 s 按 字典序排列 需要满足：对于所有有效的 i，s[i] 在字母表中的位置总是与 s[i+1] 相同或在 s[i+1] 之前。
//
// 示例 1：
//
// 输入：n = 1
// 输出：5
// 解释：仅由元音组成的 5 个字典序字符串为 ["a","e","i","o","u"]
// 示例 2：
//
// 输入：n = 2
// 输出：15
// 解释：仅由元音组成的 15 个字典序字符串为
// ["aa","ae","ai","ao","au","ee","ei","eo","eu","ii","io","iu","oo","ou","uu"]
// 注意，"ea" 不是符合题意的字符串，因为 'e' 在字母表中的位置比 'a' 靠后
// 示例 3：
//
// 输入：n = 33
// 输出：66045

func countVowelStrings(n int) int {
	// 记忆搜索,f[i][j]表示选择i个元音字符，以j结尾的字符串个数
	f := make([][5]int, n)
	var dfs func(i, j int) int
	dfs = func(i, j int) int {
		// 表示选择了n个元音字符，计数+1
		if i >= n {
			return 1
		}
		// 记忆矩阵返回，避免重复递归
		if f[i][j] != 0 {
			return f[i][j]
		}
		// 不够n个元音字符，则从j开始到结束进行选择元音字符
		cnt := 0
		for k := j; k < 5; k++ {
			cnt += dfs(i+1, k)
		}
		f[i][j] = cnt
		return f[i][j]
	}
	return dfs(0, 0)
}

func main() {
	cases := []int{
		1, 2, 33,
	}
	for _, v := range cases {
		fmt.Println(countVowelStrings(v))
	}
}
