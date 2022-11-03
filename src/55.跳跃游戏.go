package main

func canJump(nums []int) bool {
	index := len(nums) - 1
	for i := index; i >= 0; i-- {
		if nums[i]+i >= index {
			index = i
		}
	}
	return index == 0
}
