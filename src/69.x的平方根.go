package main

func mySqrt(x int) int {
	if x == 1 || x == 2 {
		return 1
	}
	if x == 0 {
		return 0
	}
	for i := 0; i < x; i++ {
		if i*i > x {
			return i - 1
		}
	}
	return 0
}

//func main() {
//	fmt.Println((2 | 24) >> 26)
//}
