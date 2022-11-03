package main

func minSwap(nums1 []int, nums2 []int) int {
	var dp [1000001][2]int
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < len(nums1); i++ {
		var a bool = nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1]
		var b bool = nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1]
		if a && b {
			dp[i][0] = min(dp[i-1][0], dp[i-1][1])
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else if a {
			dp[i][0] = dp[i-1][0]
			dp[i][1] = dp[i-1][1] + 1
		} else if b {
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		}
	}
	return min(dp[len(nums1)-1][0], dp[len(nums1)-1][1])
}
func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

//func main() {
//	nums1 := []int{0, 3, 5, 8, 9}
//	nums2 := []int{2, 1, 4, 6, 9}
//	fmt.Println(minSwap(nums1, nums2))
//}
