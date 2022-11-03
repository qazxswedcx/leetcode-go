package main

func buildArray(target []int, n int) []string {
	index := 0
	var res []string
	for i := 1; i <= n; i++ {
		if i == target[index] {
			res = append(res, "Push")
			index += 1
		} else {
			res = append(res, "Push")
			res = append(res, "Pop")
		}
		if index == len(target) {
			break
		}
	}
	return res
}

//func main() {
//	target := []int{1, 3}
//	fmt.Println(buildArray(target, 4))
//}
