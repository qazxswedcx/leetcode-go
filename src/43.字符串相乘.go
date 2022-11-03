package main

import (
	"strconv"
)

func multiply(num1 string, num2 string) string {
	add := []string{}
	for i := len(num2) - 1; i >= 0; i-- {
		a := int(num2[i] - '0')
		add1 := multiply1(num1, a)
		num0 := len(num2) - 1 - i
		for {
			if num0 == 0 {
				break
			}
			add1 += "0"
			num0--
		}
		add = append(add, add1)
	}
	res := ""
	for _, s := range add {
		res = addStrings1(res, s)
	}
	index := 0
	for i := 0; i < len(res); i++ {
		if res[i] == '0' {
			index++
		} else {
			break
		}
	}
	if index == len(res) {
		return "0"
	}
	return res[index:]
}

func multiply1(num1 string, num2 int) string {
	a, c := 0, 0
	i := len(num1) - 1
	res := ""
	for {
		if i < 0 {
			if c != 0 {
				return strconv.Itoa(c) + res
			}
			return res
		}
		a = int(num1[i] - '0')
		res = strconv.Itoa((a*num2+c)%10) + res
		c = (a*num2 + c) / 10
		i--
		a = 0
	}
}

func addStrings1(num1 string, num2 string) string {
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
//	fmt.Println(multiply("12", "1"))
//}
