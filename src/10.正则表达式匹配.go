package main

import "fmt"

func isMatch(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] != '*' {
				if i >= 1 && (p[j-1] == s[i-1] || p[j-1] == '.') {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			} else {

				if i >= 1 && j >= 2 && (p[j-2] == '.' || p[j-2] == s[i-1]) {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j] || dp[i][j-2]
				}
			}
			fmt.Println(i, j, dp[i][j])
		}

	}
	return dp[m][n]
}

//func main() {
//	fmt.Println(isMatch("aab", "c*a*b"))
//}
