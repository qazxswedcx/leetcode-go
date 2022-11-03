package main

func isMatch1(s string, p string) bool {
	m, n := len(s), len(p)
	dp := make([][]bool, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, n+1)
	}
	dp[0][0] = true
	for i := 1; i <= n; i++ {
		if p[i-1] == '*' {
			dp[0][i] = true
		} else {
			break
		}
	}
	for i := 0; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] != '*' {
				if i >= 1 && (p[j-1] == s[i-1] || p[j-1] == '?') {
					dp[i][j] = dp[i][j] || dp[i-1][j-1]
				}
			} else {
				if i >= 1 && j >= 1 {
					dp[i][j] = dp[i][j-1] || dp[i-1][j]
				}
			}
		}

	}
	return dp[m][n]
}

//func main() {
//	fmt.Println(isMatch1("baaaaaaaaabbababbbaaaabbababaabbbbabaabaaababababaabaababbbbabbbabaababaaaaaaaaaabbabaaaaabababbbbabbbbabbaabaabbbbaaaaaaaababaaaababaabbabbabaaaaaaabaababbbbbaaaabaaaaabbaababbbbbaabbaaaaaaaaaabbaababaaa",
//		"bbb*aaaaaab***ba**aab***aa*b**b*a**a**a**b*bb**a**a*a**a*b*bb*bb***b**aba**b*bbb**aab*aba**b***a*bbb*b"))
//}
