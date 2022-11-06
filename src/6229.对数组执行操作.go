package main

func applyOperations(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			nums[i] *= 2
			nums[i+1] = 0
		}
	}
	//fmt.Println(nums)
	res := make([]int, len(nums))
	index := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			res[index] = nums[i]
			index++
		}
	}
	return res
}

//func main() {
//	fmt.Println(applyOperations([]int{0, 1}))
//}
