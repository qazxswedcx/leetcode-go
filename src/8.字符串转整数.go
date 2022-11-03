package main

import (
	"strconv"
)

func myAtoi(s string) int {
	flag := false
	var start, end int
	start = 0
	end = len(s)
	op := 0
	for i := 0; i < len(s); i++ {
		if flag == false && (s[i] == '-' || s[i] == '+') {
			op++
			if op == 2 {
				return 0
			}
		}

		if flag == false && !(s[i] == ' ' || s[i] == '-' || s[i] == '+' || (s[i] >= '0' && s[i] <= '9')) {
			return 0
		}
		if flag == false && s[i] >= '0' && s[i] <= '9' {
			start = i
			flag = true
		}
		if flag == true && (s[i] < '0' || s[i] > '9') {
			end = i
			break
		}
	}
	if op == 1 && start != 0 && (s[start-1] != '+' && s[start-1] != '-') {
		return 0
	}
	if start != 0 && s[start-1] == '-' {
		start--
	}
	ret := s[start:end]

	re, _ := strconv.Atoi(ret)
	if re < -2<<30 {
		re = -2 << 30
	}
	if re > 2<<30-1 {
		re = 2<<30 - 1
	}
	return re
}

//func main() {
//	re := myAtoi("+")
//	print(re)
//}
