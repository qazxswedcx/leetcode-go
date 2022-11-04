package main

func getPermutation(n int, k int) string {
	str := "123456789"
	str = str[:n]
	res := ""
	num := 1
	for i := 1; i < n; i++ {
		num *= i
	}
	k = k - 1
	for i := n - 1; i > 0; i-- {
		b := k / num
		k = k % num
		num = num / i
		res = res + str[b:b+1]
		str = str[:b] + str[b+1:]
		//fmt.Printf("res:%v  str:%v\n", res, str)
	}
	return res + str
}

//func main() {
//	fmt.Println(getPermutation(4, 9))
//}
