package main

func removeElement(nums []int, val int) int {
	var count int = 0
	var right int = len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			count++
			for {
				if right < 0 {
					break
				}
				if nums[right] != val {
					nums[i] = nums[right]
					right--
					break
				}
				right--
			}
		}
	}
	return len(nums) - count
}

//func main() {
//	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
//	length := removeElement(nums, 2)
//	fmt.Println(nums, length)
//}
