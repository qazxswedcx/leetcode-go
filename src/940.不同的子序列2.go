package main

func distinctSubseqII(s string) int {
	dp := make([]int, 0)
	dp = append(dp, 1)
	count := 1
	for i := 1; i < len(s); i++ {
		count = (count*2 + 1) % (1e9 + 7)
		for j := 0; j < i; j++ {
			if s[i] == s[j] && j > 0 {
				count = (count - dp[j]%(1e9+7) + dp[j-1]%(1e9+7) + (1e9 + 7))
			} else if s[i] == s[j] && j == 0 {
				count = (count - dp[j]%(1e9+7) + (1e9 + 7))
			}
		}
		dp = append(dp, count%(1e9+7))
	}

	return dp[len(s)-1] % (1e9 + 7)
}

//func main() {
//
//	fmt.Println(distinctSubseqII("zchmliaqdgvwncfatcfivphddpzjkgyygueikthqzyeeiebczqbqhdytkoawkehkbizdmcnilcjjlpoeoqqoqpswtqdpvszfaksn"))
//}
