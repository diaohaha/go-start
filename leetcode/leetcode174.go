package main

import "fmt"

func calculateMinimumHP(dungeon [][]int) int {
	// 反转列表
	newdungeon := [][]int{}
	for _, row := range dungeon {
		newRow := []int{}
		for _, val := range row {
			newRow = append([]int{val}, newRow...)
		}
		newdungeon = append([][]int{newRow}, newdungeon...)
	}
	m := len(dungeon[0])
	n := len(dungeon)
	fmt.Println(newdungeon, m, n)
	return calculateMinimumHPFunc(newdungeon, n-1, m-1)
}

func calculateMinimumHPFunc(dungeon [][]int, m int, n int) int {
	if m > 0 && n > 0 {
		v1 := calculateMinimumHPFunc(dungeon, m, n-1)
		v2 := calculateMinimumHPFunc(dungeon, m-1, n)
		if v1 > v2 {
			if (v2 - dungeon[m][n]) < 1 {
				return 1
			} else {
				return v2 - dungeon[m][n]
			}
		} else {
			if (v1 - dungeon[m][n]) < 1 {
				return 1
			} else {
				return v1 - dungeon[m][n]
			}
		}
	} else if m == 0 && n > 0 {
		v := calculateMinimumHPFunc(dungeon, m, n-1)
		if (v - dungeon[m][n]) < 1 {
			return 1
		} else {
			return v - dungeon[m][n]
		}
	} else if m > 0 && n == 0 {
		v := calculateMinimumHPFunc(dungeon, m-1, n)
		if (v - dungeon[m][n]) < 1 {
			return 1
		} else {
			return v - dungeon[m][n]
		}
	} else {
		// 0,0
		if (1 - dungeon[0][0]) < 1 {
			return 1
		} else {
			return 1 - dungeon[0][0]
		}
	}
}

func runLeetcode174() {
	test := [][]int{
		[]int{-2, -3, 3},
		[]int{-5, -10, 1},
		[]int{10, 30, -5},
	}
	fmt.Println(test)
	fmt.Println(calculateMinimumHP(test))

	test2 := [][]int{
		[]int{0, 0},
	}
	fmt.Println(test2)
	fmt.Println(calculateMinimumHP(test2))
}
