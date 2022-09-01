package main

import (
	"fmt"
)

func maxArea(height []int) int {
	max := 0
	maxJ := 0
	for j := len(height) - 1; j > 0; j-- {
		if height[j] < maxJ {
			continue
		} else {
			maxJ = height[j]
		}
		maxI := 0
		for i := 0; i < j; i++ {
			if height[i] > maxI && height[i] < height[j] {
				maxI = height[i]
				if max < maxI*(j-i) {
					max = maxI * (j - i)
				}
			}
			if height[i] >= height[j] && max < height[j]*(j-i) {
				max = height[j] * (j - i)
				break
			}
		}
		fmt.Println(max)
	}
	return max
}

func main() {
	fmt.Println("miau")
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 2, 4, 3}))
	fmt.Println(maxArea([]int{1, 2, 3, 4, 5, 25, 24, 3, 4}))
}
