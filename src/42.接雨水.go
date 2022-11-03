package main

import "fmt"

func trap(height []int) int {
	var leftMax = make([]int, len(height))
	var rightMax = make([]int, len(height))
	if len(height) <= 2 {
		return 0
	}
	var res = 0

	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		if leftMax[i-1] > height[i] {
			leftMax[i] = leftMax[i-1]
		} else {
			leftMax[i] = height[i]
		}
	}

	rightMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		if rightMax[i+1] > height[i] {
			rightMax[i] = rightMax[i+1]
		} else {
			rightMax[i] = height[i]
		}
	}

	for i := 0; i < len(height); i++ {
		if leftMax[i] < rightMax[i] {
			res = res + leftMax[i] - height[i]
		} else {
			res = res + rightMax[i] - height[i]
		}
	}
	fmt.Println(leftMax, rightMax)
	return res
}

//func main() {
//	var height = []int{4, 2, 0, 3, 2, 5}
//	fmt.Println(trap(height))
//}
