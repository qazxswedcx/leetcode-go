package main

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	monoStack := []int{}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(monoStack) > 0 && monoStack[len(monoStack)-1] <= nums[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			res[i] = -1
		} else {
			res[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, nums[i])
	}
	for i := len(nums) - 1; i >= 0; i-- {
		for len(monoStack) > 0 && monoStack[len(monoStack)-1] <= nums[i] {
			monoStack = monoStack[:len(monoStack)-1]
		}
		if len(monoStack) == 0 {
			res[i] = -1
		} else {
			res[i] = monoStack[len(monoStack)-1]
		}
		monoStack = append(monoStack, nums[i])
	}
	return res
}

//func main() {
//	nums2 := []int{1, 2, 3, 4, 3}
//	fmt.Println(nextGreaterElements(nums2))
//}
