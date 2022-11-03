package main

func nextPermutation(nums []int) {
	l := len(nums)
	for i := l - 1; i >= 0; i-- {
		if i == 0 {
			reverse(nums)
			break
		}
		if nums[i] > nums[i-1] {
			reverse(nums[i:l])
			for j := i; j < l; j++ {
				if nums[j] > nums[i-1] {
					nums[j], nums[i-1] = nums[i-1], nums[j]
					break
				}
			}
			break
		}
	}
}

func reverse(a []int) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

//func main() {
//	a := []int{4, 3, 5, 6, 1, 5, 4}
//	nextPermutation(a)
//	fmt.Println(a)
//}
