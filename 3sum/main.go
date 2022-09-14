package main

import (
	"fmt"
	"sort"
)

func binarySearch(haystack []int, needle int) bool {
	low := 0
	high := len(haystack) - 1
	for low <= high {
		median := (low + high) / 2
		if haystack[median] < needle {
			low = median + 1
		} else {
			high = median - 1
		}
	}
	if low == len(haystack) || haystack[low] != needle {
		return false
	}
	return true
}

// Runtime: 102 ms, faster than 35.48% of Go online submissions for 3Sum.
// Memory Usage: 7.7 MB, less than 47.35% of Go online submissions for 3Sum.
func threeSum(nums []int) [][]int {
	var posN []int
	var negN []int
	var zeros []int
	var out [][]int
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			posN = append(posN, nums[i])
		} else if nums[i] < 0 {
			negN = append(negN, nums[i])
		} else {
			zeros = append(zeros, nums[i])
		}
	}
	sort.Ints(posN)
	sort.Ints(negN)

	prevI := 60000000 // Random numbers greater than 10^5
	prevJ := 60000000
	for i := 0; i < len(posN); i++ {
		if prevI != posN[i] {
			for j := i + 1; j < len(posN); j++ {
				if prevJ != posN[j] {
					sum := posN[i] + posN[j]
					// binary search
					found := binarySearch(negN, -sum)
					if found {
						out = append(out, []int{posN[i], posN[j], -sum})
					}
				}
				prevJ = posN[j]
			}
		}
		prevJ = 60000000
		prevI = posN[i]
	}

	prevI = 60000000
	prevJ = 60000000
	for i := 0; i < len(negN); i++ {
		if prevI != negN[i] {
			for j := i + 1; j < len(negN); j++ {
				if prevJ != negN[j] {
					sum := negN[i] + negN[j]
					// binary search
					found := binarySearch(posN, -sum)
					if found {
						out = append(out, []int{negN[i], negN[j], -sum})
					}
				}
				prevJ = negN[j]
			}
		}
		prevJ = 60000000
		prevI = negN[i]
	}

	if len(zeros) > 0 {
		prevI = 60000000
		for i := 0; i < len(posN); i++ {
			if prevI != posN[i] {
				found := binarySearch(negN, -posN[i])
				if found {
					out = append(out, []int{posN[i], -posN[i], 0})
				}
			}
			prevI = posN[i]
		}
	}

	if len(zeros) > 2 {
		out = append(out, []int{0, 0, 0})
	}

	return out
}

// Input
// [3,0,-2,-1,1,2]
// Output
// [[1,-1,0],[2,-2,0]]
// Expected
// [[-2,-1,3],[-2,0,2],[-1,0,1]]

func main() {
	fmt.Println("asdf")
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4, 5}))
	fmt.Println(threeSum([]int{3, 0, -2, -1, 1, 2}))
	//[[-2,-1,3],[-2,0,2],[-1,0,1]]

}
