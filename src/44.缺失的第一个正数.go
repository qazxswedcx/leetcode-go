package main

func firstMissingPositive(nums []int) int {
	length := len(nums)
	for i := 0; i < length; i++ {
		if nums[i] <= 0 {
			nums[i] = length + 1
		}
	}
	for i := 0; i < length; i++ {
		num := abs(nums[i])
		if num <= length {
			nums[num-1] = -abs(nums[num-1])
		}
	}
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return length + 1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//func main() {
//	nums := []int{3, 4, -1, 1}
//	fmt.Println(firstMissingPositive(nums))
//}

//func firstMissingPositive(nums []int) int {
//	sort.Ints(nums)
//	flag := false
//	count := 1
//	for i := 0; i < len(nums); i++ {
//		if nums[i] > 0 {
//			flag = true
//		}
//		if flag == true && nums[i] == count {
//			count++
//		} else if flag == true && nums[i] > count {
//			return count
//		}
//	}
//	return count
//}
