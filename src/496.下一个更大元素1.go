package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	res := map[int]int{}
	monostack := []int{}
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(monostack) > 0 && monostack[len(monostack)-1] < nums2[i] {
			monostack = monostack[:len(monostack)-1]
		}
		if len(monostack) == 0 {
			res[nums2[i]] = -1
		} else {
			res[nums2[i]] = monostack[len(monostack)-1]
		}
		monostack = append(monostack, nums2[i])
	}
	result := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = res[nums1[i]]
	}
	return result
}

//func main() {
//	nums1 := []int{4, 1, 2}
//	nums2 := []int{1, 3, 4, 2}
//	fmt.Println(nextGreaterElement(nums1, nums2))
//}
