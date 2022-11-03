package main

func sortColors(nums []int) {
	a := []int{0, 0, 0}
	for i := 0; i < len(nums); i++ {
		a[nums[i]]++
	}
	for i := 0; i < len(nums); i++ {
		if i < a[0] {
			nums[i] = 0
		} else if i >= a[0] && i < a[0]+a[1] {
			nums[i] = 1
		} else {
			nums[i] = 2
		}
	}
}
