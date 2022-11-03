package main

func maxWidthRamp(nums []int) int {
	res := 0
	monoStack := []int{}
	for i := 0; i < len(nums); i++ {
		if len(monoStack) == 0 || nums[monoStack[len(monoStack)-1]] > nums[i] {
			monoStack = append(monoStack, i)
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(monoStack) > 0 && nums[i] >= nums[monoStack[len(monoStack)-1]] {
			res = max(res, i-monoStack[len(monoStack)-1])
			monoStack = monoStack[:len(monoStack)-1]
		}
	}
	return res
}

//func main() {
//	nums := []int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}
//	fmt.Println(maxWidthRamp(nums))
//}
