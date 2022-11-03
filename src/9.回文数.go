package main

import (
	"fmt"
	"strconv"
)

// 转化字符串方法
func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	for i := 0; i < len(str); i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}

// 数值方法
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	n := x
	newed := 0
	for {

		newed = newed*10 + x%10
		x = x / 10
		if x < 1 {
			break
		}
	}
	fmt.Println(newed)
	if newed == n {
		return true
	} else {
		return false
	}
}

//func main() {
//	fmt.Println(isPalindrome2(1233321))
//}
