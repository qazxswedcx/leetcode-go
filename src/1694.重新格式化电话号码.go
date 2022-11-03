package main

import (
	"strings"
)

func reformatNumber(number string) string {
	number = strings.ReplaceAll(number, "-", "")
	number = strings.ReplaceAll(number, " ", "")
	length := len(number)
	res := ""
	for i := 0; i < length; i++ {
		res = res + number[i:i+1]
		if i%3 == 2 && i != length-1 {
			res = res + "-"
		}

	}
	if len(res)%4 == 1 {
		res = res[:len(res)-2] + res[len(res)-1:]
		res = res[:len(res)-2] + "-" + res[len(res)-2:]
	}
	return res
}

//func main() {
//	number := "--17-5 229 35-39475 "
//	fmt.Println(reformatNumber(number))
//}
