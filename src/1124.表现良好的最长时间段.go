package main

func longestWPI(hours []int) int {
	res := 0
	prefixSum := make([]int, len(hours)+1)
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			hours[i] = 1
		} else {
			hours[i] = -1
		}
		prefixSum[i+1] = prefixSum[i] + hours[i]
	}
	monoStack := []int{}
	for i := 0; i < len(prefixSum); i++ {
		if len(monoStack) == 0 || prefixSum[monoStack[len(monoStack)-1]] > prefixSum[i] {
			monoStack = append(monoStack, i)
		}
	}
	for i := len(prefixSum) - 1; i >= 0; i-- {
		for len(monoStack) > 0 && prefixSum[i] > prefixSum[monoStack[len(monoStack)-1]] {
			res = max(res, i-monoStack[len(monoStack)-1])
			monoStack = monoStack[:len(monoStack)-1]
		}
	}
	return res
}

//func main() {
//	hours := []int{11, 2, 4, 14, 2, 15, 7, 10, 1, 16, 9, 0, 2, 8, 4, 14, 6, 12, 2, 8, 6, 4, 14, 13, 7, 16, 14, 2, 3, 2, 8, 3, 12, 3, 3, 9, 14, 1, 5, 3, 12, 0, 15, 5, 0, 2, 3, 16, 7, 2, 1, 1, 4, 9, 0, 11, 9, 16, 15, 7, 0, 5, 6, 4, 12, 1, 1, 2, 13, 8, 3, 9, 12, 9, 3, 11, 4, 14, 7, 5, 16, 0, 11, 8, 8, 14, 1, 5, 0, 6, 5, 8, 10, 15, 9, 14, 16, 11, 1, 13}
//	fmt.Println(longestWPI(hours))
//}
