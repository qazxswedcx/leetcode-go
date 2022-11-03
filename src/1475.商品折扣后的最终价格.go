package main

func finalPrices(prices []int) []int {
	res := make([]int, len(prices))
	stk := []int{}
	for i := len(prices) - 1; i >= 0; i-- {
		for len(stk) > 0 && stk[len(stk)-1] > prices[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			res[i] = prices[i]
		} else {
			res[i] = prices[i] - stk[len(stk)-1]
		}
		stk = append(stk, prices[i])
	}
	return res
}

//func main() {
//	fmt.Println(finalPrices([]int{8, 4, 6, 2, 3}))
//}
