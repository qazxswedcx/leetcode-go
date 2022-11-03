package main

func isValidSudoku(board [][]byte) bool {
	var line, column [9][9]bool //行、列
	var block [3][3][9]bool     //九宫格
	for i, row := range board {
		for j, b := range row {
			if b != '.' {
				digit := b - '1'
				if line[i][digit] || column[j][digit] || block[i/3][j/3][digit] {
					return false
				}
				line[i][digit] = true         //第i行有数字digit
				column[j][digit] = true       //第i列有数字digit
				block[i/3][j/3][digit] = true //第i九宫格有数字digit
			}
		}
	}
	return true
}

//func main() {
//	board := [][]byte{
//		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
//		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
//		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
//		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
//		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
//		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
//		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
//		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
//		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
//	}
//	fmt.Println(isValidSudoku(board))
//}
