package main

func maxArea(height []int) int {
	res := 0
	leng := len(height)
	left, right := 0, leng-1
	for {
		if left == right {
			break
		}
		if height[left] < height[right] {
			if res < (right-left)*height[left] {
				res = (right - left) * height[left]
			}
			left++
		} else {
			if res < (right-left)*height[right] {
				res = (right - left) * height[right]
			}
			right--
		}

	}

	return res
}

//func main() {
//	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
//	fmt.Println(maxArea(a))
//}
