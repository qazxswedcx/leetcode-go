package main

func sumSubarrayMins(arr []int) (ans int) {
	const mod int = 1e9 + 7
	n := len(arr)
	monoStack := []int{}
	dp := make([]int, n)
	for i, x := range arr {
		for len(monoStack) > 0 && arr[monoStack[len(monoStack)-1]] > x {
			monoStack = monoStack[:len(monoStack)-1] //前一个数大于当前数
		}
		k := i + 1
		if len(monoStack) > 0 {
			k = i - monoStack[len(monoStack)-1]
		}
		dp[i] = k * x
		if len(monoStack) > 0 {
			dp[i] += dp[i-k]
		}
		ans = (ans + dp[i]) % mod
		monoStack = append(monoStack, i)
	}
	return
}

//func main() {
//	arr := []int{11, 81, 94, 43, 3}
//	fmt.Println(sumSubarrayMins(arr))
//}

//func sumSubarrayMins(arr []int) int {
//	minNum := make([]int, len(arr)+1)
//	res := 0
//	for i := 0; i < len(arr); i++ {
//		for j := i; j >= 0; j-- {
//			if j == i {
//				minNum[j] = arr[i]
//				res += minNum[j]
//			} else {
//				if arr[i] < minNum[j] {
//					minNum[j] = arr[i]
//				}
//
//				res += minNum[j]
//			}
//		}
//	}
//	return res % (1e9 + 7)
//}
