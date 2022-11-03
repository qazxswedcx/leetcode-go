package main

func removeDuplicates(nums []int) int {
	var a int = 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[a] {
			a += 1
			nums[a] = nums[i]
		}
	}
	return a + 1
}

//func main() {
//	a := []int{1, 2, 2, 3, 4, 4, 5, 5}
//	fmt.Println(removeDuplicates(a), a)
//
//}
