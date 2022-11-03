package main

func possibleBipartition(n int, dislikes [][]int) bool {
	g := make([][]int, n)
	for _, d := range dislikes {
		x, y := d[0]-1, d[1]-1
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
	}
	color := make([]int, n) // color[x] = 0 表示未访问节点 x
	var dfs func(int, int) bool
	dfs = func(x, c int) bool {
		color[x] = c
		for _, y := range g[x] {
			if color[y] == c || color[y] == 0 && !dfs(y, 3^c) {
				return false
			}
		}
		return true
	}
	for i, c := range color {
		if c == 0 && !dfs(i, 1) {
			return false
		}
	}
	return true
}

//func main() {
//	dislikes := [][]int{{1, 2}, {1, 3}, {2, 4}}
//	fmt.Println(possibleBipartition(4, dislikes))
//}
