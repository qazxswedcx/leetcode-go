package main

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	preSumArr := make([]int, n+1)
	for i, num := range nums {
		preSumArr[i+1] = preSumArr[i] + num
	}
	ans := n + 1
	q := []int{}
	for i, curSum := range preSumArr {
		for len(q) > 0 && curSum-preSumArr[q[0]] >= k {
			ans = min(ans, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && preSumArr[q[len(q)-1]] >= curSum {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if ans < n+1 {
		return ans
	}
	return -1
}

//func main() {
//	nums := []int{2, -1, 2}
//	fmt.Println(shortestSubarray(nums, 3))
//}
