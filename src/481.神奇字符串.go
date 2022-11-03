package main

import "fmt"

func magicalString(n int) int {
	s := "122"
	if n <= 3 {
		return 1
	}
	index := 2
	nums := 1
	for i := 3; i < n; {
		if s[index] == '2' {
			if i == n-1 && s[i-1] == '2' {
				return nums + 1
			}
			if s[i-1] == '2' {
				s = s + "11"
				nums += 2
			} else {
				s = s + "22"
			}
			i += 2
		} else {
			if s[i-1] == '2' {
				s = s + "1"
				nums += 1
			} else {
				s = s + "2"
			}
			i += 1
		}
		index++
	}
	fmt.Println(s)
	return nums
}

//func main() {
//	fmt.Println(magicalString(4))
//}
