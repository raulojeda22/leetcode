package main

import (
	"fmt"
	"strconv"
)

/**
*	This was my first aproach but, it was too slow
 */
func lengthOfLongestSubstring(s string) int {
	var r int = 0
	length := len(s)
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			subs := s[i : j+1]
			lS := len(subs)
			if lS > r {
				a := [255]bool{}
				inter := false
				for _, v := range subs {
					if a[v] {
						inter = true
						break
					} else {
						a[v] = true
					}
				}
				if !inter && r < lS {
					r = lS
				}
			}
		}
	}
	return r
}

/**
*	Think of a better way to calculate the answer instead of jumping straight to the brute force
 */
func lengthOfLongestSubstring2(s string) int {
	var r int = 0
	length := len(s)
	for i := 0; i < length; i++ {
		a := [255]bool{}
		for j := i; j < length; j++ {
			c := s[j]
			if a[c] {
				break
			} else {
				o := (j + 1 - i)
				if r < o {
					r = o
				}
				a[c] = true
			}
		}
	}
	return r
}

func main() {
	a := []string{
		"",
		" ",
		"asdf",
		"aaab",
		"qweraaa",
		"abcabcbb",
		"pwwkew",
	}
	for _, v := range a {
		fmt.Println(v + " " + strconv.Itoa(lengthOfLongestSubstring2(v)))
	}
}
