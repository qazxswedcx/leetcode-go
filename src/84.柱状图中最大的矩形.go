package main

func largestRectangleArea(heights []int) int {
	maxArea := 0
	monoStack := []int{}
	for i, height := range heights {
		for len(monoStack) > 0 && heights[monoStack[len(monoStack)-1]] > height {
			h := heights[monoStack[len(monoStack)-1]] //被弹出的高度
			monoStack = monoStack[:len(monoStack)-1]
			if len(monoStack) == 0 { //弹出之后没有值，说明前面的都大于该数，宽为该数
				maxArea = max(maxArea, h*i)
			} else {
				maxArea = max(maxArea, h*(i-monoStack[len(monoStack)-1]-1))
			}

		}
		monoStack = append(monoStack, i)
		if i == len(heights)-1 {
			for len(monoStack) > 0 {
				h := heights[monoStack[len(monoStack)-1]]
				monoStack = monoStack[:len(monoStack)-1]
				if len(monoStack) == 0 { //弹出之后没有值，说明前面的都大于该数，宽为该数
					maxArea = max(maxArea, h*len(heights))
				} else {
					maxArea = max(maxArea, h*(len(heights)-monoStack[len(monoStack)-1]-1))
				}
			}
		}
	}
	return maxArea
}

//func main() {
//	height := []int{2, 4}
//	fmt.Println(largestRectangleArea(height))
//}
