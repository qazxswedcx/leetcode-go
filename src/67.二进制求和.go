package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	m, n := len(a), len(b)
	if m > n {
		for i := 0; i < m-n; i++ {
			b = "0" + b
		}
	} else if m < n {
		for i := 0; i < n-m; i++ {
			a = "0" + a
		}
	}
	res := ""
	a1, b1, c := 0, 0, 0
	for i := len(a) - 1; i >= 0; i-- {
		a1 = int(a[i] - '0')
		b1 = int(b[i] - '0')
		if a1+b1+c < 2 {
			res = strconv.Itoa(a1+b1+c) + res
			c = 0
		} else {
			res = strconv.Itoa(a1+b1+c-2) + res
			c = 1
		}
	}
	if c == 1 {
		return "1" + res
	}
	return res
}

//func main() {
//	fmt.Println(addBinary("1010", "1011"))
//}
