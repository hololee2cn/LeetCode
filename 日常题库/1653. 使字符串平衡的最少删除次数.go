package main

import "fmt"

func minimumDeletions(s string) int {
	// 如果前i个是a则无须删除
	// 统计前i个位置的b的数量，当第i个是a，则两种选择，删除a或者删除前i个b
	min := func(v1, v2 int) int {
		if v1 < v2 {
			return v1
		}
		return v2
	}
	cntB := 0
	res := 0
	for _, c := range s {
		if c == 'b' {
			cntB++
		} else {
			res = min(res+1, cntB)
		}
	}
	return res
}

func main() {
	cases := []string{
		"aababbab",
		"bbaaaaabb",
	}
	for _, v := range cases {
		fmt.Println(minimumDeletions(v))
	}
}
