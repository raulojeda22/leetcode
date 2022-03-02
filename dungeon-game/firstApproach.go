package main

import (
	"fmt"
	"math"
)

/**
 * I like this attempt so I'm going to keep it, but it does not work
 * In this attempt I tried to create 2 bidimensional arrays, one for storing the current sum
 * and other for storing the worst case until now. The problem is that you can't never really
 * know which is going to be the worst case and which addition should be stored, since for
 * example, in this array [[1, -3, 3],[0, -2, 0]] the worst case would be [[1 -2 -2][1 -1 -1]]
 * but for [[1, -3, 3][0, -2, 0][-3, -3, -3]] it should be [[1 -2 -2][1 -1 -2][-2 -5 -2]]
 * Notice that -1 changes to -2 after adding a new row, which means that we can't really tell
 * the worst case without looking at the path instead of just choosing the best cells till that
 * point like this function does. Good luck with the other approach future me, love you <3
 **/
func calculateMinimumHP1(dungeon [][]int) int {
	added := [][]int{}
	worst := [][]int{}
	for i := 0; i < len(dungeon); i++ {
		added = append(added, make([]int, len(dungeon[i])))
		for j := 0; j < len(dungeon[i]); j++ {
			added[i][j] = -4000000
		}
	}
	for i := 0; i < len(dungeon); i++ {
		worst = append(worst, make([]int, len(dungeon[i])))
		for j := 0; j < len(dungeon[i]); j++ {
			worst[i][j] = -4000000
		}
	}
	added[0][0] = dungeon[0][0]
	worst[0][0] = dungeon[0][0]
	for i := 0; i < len(dungeon); i++ {
		for j := 0; j < len(dungeon[i]); j++ {
			for k := 0; k < 2; k++ { // k is for looking down (0) or right (1)
				if k == 0 { // down
					if i+1 < len(dungeon) {
						if worst[i][j] > worst[i+1][j] && added[i][j] > added[i+1][j] {
							added[i+1][j] = added[i][j] + dungeon[i+1][j]
							if dungeon[i+1][j] >= 0 || worst[i][j] <= added[i+1][j] {
								worst[i+1][j] = worst[i][j]
							} else if worst[i][j] > added[i+1][j] {
								worst[i+1][j] = added[i+1][j]
							}
						}
					}
				} else { // right
					if j+1 < len(dungeon[i]) {
						if worst[i][j] > worst[i][j+1] && added[i][j] > added[i][j+1] {
							added[i][j+1] = added[i][j] + dungeon[i][j+1]
							if dungeon[i][j+1] >= 0 || worst[i][j] <= added[i][j+1] {
								worst[i][j+1] = worst[i][j]
							} else if worst[i][j] > added[i][j+1] {
								worst[i][j+1] = added[i][j+1]
							}
						}
					}
				}
			}
		}
	}
	for i := 0; i < len(dungeon); i++ {
		fmt.Println(dungeon[i])
	}
	for i := 0; i < len(added); i++ {
		fmt.Println(added[i])
	}
	for i := 0; i < len(worst); i++ {
		fmt.Println(worst[i])
	}
	r := 0
	v := worst[len(worst)-1][len(worst[len(worst)-1])-1]
	if v > 0 {
		r = 1
	} else {
		r = int(math.Abs(float64(v))) + 1
	}
	return r
}

func firstApproach() {
	a := [][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}
	a = [][]int{{0}}
	a = [][]int{{2, 1}, {1, -1}} // 1
	a = [][]int{{-3}, {5}}
	a = [][]int{{100}}
	a = [][]int{{-200}}
	a = [][]int{{0, 0}, {-5, -4}}
	a = [][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}
	a = [][]int{{3, -20, 30}, {-3, 4, 0}}
	fmt.Println(calculateMinimumHP1(a))
}
