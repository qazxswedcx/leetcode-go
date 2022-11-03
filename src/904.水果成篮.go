package main

func totalFruit(fruits []int) int {
	cnt := map[int]int{}
	ans, left := 0, 0
	for right, fruit := range fruits {
		cnt[fruit]++
		if len(cnt) > 2 {
			y := fruits[left]
			cnt[y]--
			if cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		if ans < right-left+1 {
			ans = right - left + 1
		}
	}
	return ans
}

//func main() {
//	fruit := []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}
//	fmt.Println(totalFruit(fruit))
//}
