package main

import "fmt"

func isPalindrome(x int) bool {
	var digits []int
	y := x
	if x < 0 {
		return false
	}
	for y > 0 {
		digits = append(digits, y%10)
		y = y / 10
	}
	equal := true
	for i := 0; i < len(digits) && i != len(digits)-i-1; i++ {
		if digits[i] != digits[len(digits)-i-1] {
			equal = false
			break
		}
	}
	return equal
}

func main() {
	fmt.Println(isPalindrome(-1123211))
}
