package main

func maxChunksToSorted(arr []int) int {
	count := 0
	maxflag := 0
	for i := 0; i < len(arr); i++ {
		if maxflag < arr[i] {
			maxflag = arr[i]
		}
		if maxflag <= i {
			count += 1
		}
	}
	return count
}

//func main() {
//	a := []int{1, 0, 2, 3, 4}
//	fmt.Println(maxChunksToSorted(a))
//}
