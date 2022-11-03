package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m + n - 1; i >= 0; i-- {
		if n == 0 {
			break
		}
		if m == 0 {
			for j := n - 1; j >= 0; j-- {
				nums1[j] = nums2[j]
			}
			break
		}
		if nums1[m-1] >= nums2[n-1] {
			nums1[i] = nums1[m-1]
			m--
		} else {
			nums1[i] = nums2[n-1]
			n--
		}
	}

}

//func main() {
//	nums1 := []int{2, 2, 3, 0, 0, 0}
//	nums2 := []int{1, 5, 6}
//	merge(nums1, 3, nums2, 3)
//	fmt.Println(nums1)
//}

/*func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := m; i < m+n; i++ {
		nums1[i] = nums2[i-m]
	}
	sort.Ints(nums1)
}*/
