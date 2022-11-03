package main

func solveSudoku(board [][]byte) {
	var line, column [9][9]bool //行、列
	var block [3][3][9]bool     //九宫格
	var spaces [][2]int         //空白格

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				spaces = append(spaces, [2]int{i, j}) //所有空白格坐标
			} else {
				digit := b - '1'
				line[i][digit] = true         //第i行有数字digit
				column[j][digit] = true       //第i列有数字digit
				block[i/3][j/3][digit] = true //第i九宫格有数字digit
			}
		}
	}

	var dfs func(int) bool
	dfs = func(pos int) bool {
		if pos == len(spaces) {
			//所有空格填满结束
			return true
		}
		i, j := spaces[pos][0], spaces[pos][1] //第pos个空格的位置
		for digit := byte(0); digit < 9; digit++ {
			if !line[i][digit] && !column[j][digit] && !block[i/3][j/3][digit] {
				//行列九宫格都没有则填入
				line[i][digit] = true
				column[j][digit] = true
				block[i/3][j/3][digit] = true
				board[i][j] = digit + '1'
				if dfs(pos + 1) {
					return true
				}
				line[i][digit] = false
				column[j][digit] = false
				block[i/3][j/3][digit] = false
			}
		}
		return false
	}
	dfs(0)
}

//func main() {
//	board := [][]byte{
//		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
//		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
//		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
//		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
//		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
//		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
//		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
//		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
//		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
//	}
//	solveSudoku(board)
//	for i := 0; i < len(board); i++ {
//		fmt.Printf("%c\n", board[i])
//	}
//
//}
