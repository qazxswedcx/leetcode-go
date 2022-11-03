package main

import (
	"math"
)

func bestCoordinate(towers [][]int, radius int) []int {
	strength := make([][]int, 51)
	for i := 0; i < len(strength); i++ {
		strength[i] = make([]int, 51)
	}
	for i := 0; i < 51; i++ {
		for j := 0; j < 51; j++ {
			for _, tower := range towers {
				strength[i][j] += calc(i, j, tower, radius)
			}

		}
	}
	x := 0
	y := 0
	for i := 0; i < len(strength); i++ {
		for j := 0; j < len(strength); j++ {
			if strength[i][j] > strength[x][y] {
				x = i
				y = j
			}
		}
	}
	return []int{x, y}
}

func calc(i0 int, i1 int, j []int, radius int) int {
	d := math.Sqrt(float64((i0-j[0])*(i0-j[0]) + (i1-j[1])*(i1-j[1])))
	if d > float64(radius) {
		return 0
	}
	q := float64(j[2]) / (1.0 + d)
	return int(q)
}

//func main() {
//	towers := [][]int{{1, 2, 5}, {2, 1, 7}, {3, 1, 9}}
//	fmt.Println(bestCoordinate(towers, 2))
//}
