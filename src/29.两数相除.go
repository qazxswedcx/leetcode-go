package main

import (
	"math"
)

func divide(dividend int, divisor int) int {
	flag := true
	if dividend > 0 && divisor < 0 || dividend < 0 && divisor > 0 {
		flag = false
	}
	if divisor == -1 && dividend == -2147483648 {
		return 2147483647
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}

	result := 0
	for i := 31; i >= 0; i-- {
		if (dividend>>i)-divisor >= 0 {
			dividend = dividend - (divisor << i)
			// 代码优化：这里控制 ans 大于等于 INT_MAX
			if result > math.MaxInt32-(1<<i) {
				return math.MinInt32
			}
			result += 1 << i
		}
	}
	if flag == true {
		return result
	} else {
		return -result
	}
}

//func main() {
//	fmt.Println(divide(-2147483648, -1))
//}
