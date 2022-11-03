package main

import (
	"strconv"
)

func countAndSay(n int) string {
	res := "1"
	for i := 1; i < n; i++ {
		res = say(res)
	}
	return res
}

func say(s string) string {
	res := ""
	temp := s[0]
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == temp {

			count += 1
		} else {

			res = res + strconv.Itoa(count) + string(temp)
			count = 1
			temp = s[i]
		}
		if i == len(s)-1 {
			res = res + strconv.Itoa(count) + string(temp)
		}
	}
	return res
}

//
//func main() {
//	fmt.Println(countAndSay(1))
//}
