package main

func shortestBridge(grid [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	n := len(grid)
	step := 0
	for i, row := range grid {
		for j, v := range row {
			if v != 1 {
				continue
			}
			island := []pair{}
			grid[i][j] = 2
			p := []pair{{i, j}}
			for len(p) > 0 {
				q := p[0]
				p = p[1:]
				island = append(island, q)
				for _, dir := range dirs {
					x, y := q.x+dir.x, q.y+dir.y
					if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
						grid[x][y] = 2
						p = append(p, pair{x, y})
					}
				}
			}
			p = island
			for {
				tmp := p
				p = nil
				for _, q := range tmp {
					for _, dir := range dirs {
						x, y := q.x+dir.x, q.y+dir.y
						if 0 <= x && x < n && 0 <= y && y < n {
							if grid[x][y] == 1 {
								return step
							}
							if grid[x][y] == 0 {
								grid[x][y] = 2
								p = append(p, pair{x, y})
							}

						}
					}
				}
				step++
			}
		}
	}
	return step
}

//func main() {
//	grid := [][]int{{1, 0}, {0, 1}}
//	fmt.Println(shortestBridge(grid))
//}

/*
func shortestBridge(grid [][]int) int {
	type pair struct{ x, y int }
	dirs := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} //上下左右四个方向
	n := len(grid)                                   //行数
	step := 0
	for i, row := range grid { //行
		for j, v := range row { //列
			if v != 1 {
				continue //寻找第一个为 1 的点
			}
			//该点是岛屿的情况
			island := []pair{}
			grid[i][j] = -1 //将值变成-1
			q := []pair{{i, j}}
			//扫描出第一个点所在岛屿的所有值
			for len(q) > 0 {
				p := q[0]                  //第一个xy
				q = q[1:]                  //剩下的值
				island = append(island, p) // 将第一个值存入岛屿中
				for _, d := range dirs {
					x, y := p.x+d.x, p.y+d.y //该点的上下左右
					if 0 <= x && x < n && 0 <= y && y < n && grid[x][y] == 1 {
						grid[x][y] = -1           //如果是就将值变成-1
						q = append(q, pair{x, y}) //该点加入岛屿
					}
				}
			}
			q = island //第一个岛屿中的所有点
			for {
				tmp := q
				q = nil
				for _, p := range tmp { //对每一个点进行扩展
					for _, d := range dirs {
						x, y := p.x+d.x, p.y+d.y
						if 0 <= x && x < n && 0 <= y && y < n {
							if grid[x][y] == 1 {
								return step //检测到岛屿结束
							}
							if grid[x][y] == 0 {
								grid[x][y] = -1 //没检测到加入
								q = append(q, pair{x, y})
							}
						}
					}
				}
				step++
			}
		}
	}
	return step
}
*/
