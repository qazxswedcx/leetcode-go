package main

func setZeroes(matrix [][]int) {
	var l1, l2 []int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				l1 = append(l1, i)
				l2 = append(l2, j)
			}
		}
	}
	for i := 0; i < len(l1); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[l1[i]][j] = 0
		}
	}
	for j := 0; j < len(l2); j++ {
		for i := 0; i < len(matrix); i++ {
			matrix[i][l2[j]] = 0
		}
	}
}

//func main() {
//	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
//	setZeroes(matrix)
//	fmt.Println(matrix)
//}
