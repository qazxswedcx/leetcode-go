package main

func romanToInt(s string) int {
	var res int = 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'M':
			res += 1000
		case 'D':
			res += 500
		case 'C':
			if i+1 < len(s) && s[i+1] == 'M' {
				res += 900
				i++
			} else if i+1 < len(s) && s[i+1] == 'D' {
				res += 400
				i++
			} else {
				res += 100
			}
		case 'L':
			res += 50
		case 'X':
			if i+1 < len(s) && s[i+1] == 'C' {
				res += 90
				i++
			} else if i+1 < len(s) && s[i+1] == 'L' {
				res += 40
				i++
			} else {
				res += 10
			}
		case 'V':
			res += 5
		case 'I':
			if i+1 < len(s) && s[i+1] == 'X' {
				res += 9
				i++
			} else if i+1 < len(s) && s[i+1] == 'V' {
				res += 4
				i++
			} else {
				res += 1
			}
		}
	}
	return res
}
func romanToInt1(s string) int {
	var romans = map[byte]int{'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000}
	var res int = 0
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && romans[s[i]] < romans[s[i+1]] {
			res -= romans[s[i]]
		} else {
			res += romans[s[i]]
		}
	}
	return res
}

//
//func main() {
//	fmt.Println(romanToInt1("MCMXCIV"))
//}
