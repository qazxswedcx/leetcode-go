package main

func searchInsert(nums []int, target int) int {
	left, right, mid := 0, len(nums), len(nums)/2
	for {

		if left == right-1 {
			if target <= nums[left] {
				return left
			} else {
				return right
			}
		}
		if target > nums[mid] {
			left = mid
			mid = (left + right) / 2
		} else if target < nums[mid] {
			right = mid
			mid = (left + right) / 2
		} else if target == nums[mid] {
			return mid
		}
	}
}

//func main() {
//	nums := []int{1}
//	fmt.Println(searchInsert(nums, 1))
//}
