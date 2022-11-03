package main

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	a, b, c := 0, 0, 0
	i := len(num1) - 1
	j := len(num2) - 1
	res := ""
	for {
		if i < 0 && j < 0 {
			if c == 1 {
				return "1" + res
			}
			return res
		}
		if i >= 0 {
			a = int(num1[i] - '0')
		}
		if j >= 0 {
			b = int(num2[j] - '0')
		}
		res = strconv.Itoa((a+b+c)%10) + res
		c = (a + b + c) / 10
		i--
		j--
		a = 0
		b = 0
	}
}

//func main() {
//	fmt.Println(addStrings("6994", "36"))
//}
