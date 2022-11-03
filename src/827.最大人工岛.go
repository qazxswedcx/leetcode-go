package main

func largestIsland(grid [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	step := 0
	index := 2
	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}

			island := []pair{}
			grid[i][j] = index
			p := []pair{{i, j}}
			for len(p) > 0 {
				q := p[0]
				p = p[1:]
				island = append(island, q)
				for _, dir := range dirs {
					x, y := q.x+dir.x, q.y+dir.y
					if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
						grid[x][y] = index
						p = append(p, pair{x, y})
					}
				}
			}
			index++
		}
	}
	nums := make([]int, index+1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] != 0 {
				nums[grid[i][j]]++
			}
		}
	}
	full := true
	for i, row := range grid {
		for j, v := range row {
			if v != 0 {
				continue
			}
			full = false
			temp := []int{}
			for _, dir := range dirs {
				x, y := i+dir.x, j+dir.y
				if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] != 0 {
					temp = append(temp, grid[x][y])
				}
			}
			temp = removeRepeatElement(temp)
			res := 1
			for i := 0; i < len(temp); i++ {
				res += nums[temp[i]]
			}
			step = max(step, res)
		}
	}
	if full == true {
		return n * n
	}
	return step
}

func removeRepeatElement(list []int) []int {
	// 创建一个临时map用来存储数组元素
	temp := make(map[int]struct{})
	index := 0
	// 将元素放入map中
	for _, v := range list {
		temp[v] = struct{}{}
	}
	tempList := make([]int, len(temp))
	for key := range temp {
		tempList[index] = key
		index++
	}
	return tempList
}

//func main() {
//	grid := [][]int{{1, 1}, {0, 1}}
//	fmt.Println(largestIsland(grid))
//}
