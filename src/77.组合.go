package main

func combine(n int, k int) (ans [][]int) {
	temp := []int{}
	var dfs func(int)
	dfs = func(cur int) {
		if len(temp)+(n-cur+1) < k {
			return
		}
		if len(temp) == k {
			comb := make([]int, k)
			copy(comb, temp)
			ans = append(ans, comb)
			return
		}
		temp = append(temp, cur)
		dfs(cur + 1)
		temp = temp[:len(temp)-1]
		dfs(cur + 1)
	}
	dfs(1)
	return
}
