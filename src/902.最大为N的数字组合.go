package main

import (
	"fmt"
	"strconv"
)

func atMostNGivenDigitSet(digits []string, n int) int {
	num := strconv.Itoa(n)
	leng := len(digits)
	count := 0
	for i := 1; i < len(num); i++ {
		count = count + pow(leng, i)
	}
	low := []int{}
	equal := []int{}
	for i := 0; i < len(num); i++ {
		low = append(low, 0)
		equal = append(equal, 0)
		for j := 0; j < leng; j++ {
			if digits[j][0] < num[i] {
				low[i]++
			}
			if digits[j][0] == num[i] {
				equal[i]++
			}
		}
	}

	if equal[0] == 0 {
		count = count + low[0]*pow(leng, len(num)-1)
	} else {
		flag := 1
		for i := 1; i < len(num); i++ {
			if equal[i] == 0 {
				flag = i
				break
			}
			if i == len(num)-1 {
				flag = len(num)
				break
			}
		}
		fmt.Println(flag)
		for i := 0; i <= flag && i < len(num); i++ {
			count = count + low[i]*pow(leng, len(num)-i-1)
		}
		if flag == len(num) {
			count += 1
		}
	}

	return count
}
func pow(a int, b int) int {
	res := 1
	for i := 0; i < b; i++ {
		res = res * a
	}
	return res
}

//func main() {
//	digits := []string{"3", "4", "5", "6"}
//	fmt.Println(atMostNGivenDigitSet(digits, 64))
//}
