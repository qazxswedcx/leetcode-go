package main

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stk := []int{}
	for i := len(temperatures) - 1; i >= 0; i-- {
		for len(stk) > 0 && temperatures[stk[len(stk)-1]] <= temperatures[i] {
			stk = stk[:len(stk)-1]
		}
		if len(stk) == 0 {
			res[i] = 0
		} else {
			res[i] = stk[len(stk)-1] - i
		}
		stk = append(stk, i)
	}
	return res
}

//func main() {
//	fmt.Println(dailyTemperatures([]int{89, 62, 70, 58, 47, 47, 46, 76, 100, 70}))
//}
