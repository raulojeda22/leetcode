package main

import "fmt"

type data struct {
	init string
	outp string
	rows int
}

/**
* This is actually pretty slow 95% of the users made it faster, so I'm going to try a faster solution
* Runtime: 816 ms, 8 MB of memory usage
 */
func convert(s string, numRows int) string {
	r := ""
	lS := len(s)
	m := [1000][1000]byte{}
	down := true
	x := 0
	y := 0
	for i := 0; i < lS; i++ {
		if down {
			m[x][y] = s[i]
			y++
			if y+1 == numRows {
				down = false
			}
		} else {
			m[x][y] = s[i]
			x++
			y--
			if y == 0 {
				down = true
			}
		}
	}
	lM := len(m)
	for i := 0; i < lM; i++ {
		for j := 0; j < len(m[i]); j++ {
			d := m[j][i]
			if d != 0 {
				r += string([]byte{d})
			}
		}
	}
	return r
}

/**
* Proud of this 😎
* Runtime: 8 ms, faster than 80.00% of Go online submissions for Zigzag Conversion.
* Memory Usage: 7 MB, less than 53.97% of Go online submissions for Zigzag Conversion.
*
* Runtime in this platform is quite variable, after changing len(s) for lS, one time it was executed in 40 ms,
* and then the same code was executed in 4 ms, it does not make much sense.
* I'm going to record that runtime because it's pretty good
*
* Runtime: 4 ms, faster than 94.76% of Go online submissions for Zigzag Conversion.
* Memory Usage: 7.1 MB, less than 53.97% of Go online submissions for Zigzag Conversion.
 */
func convert2(s string, numRows int) string {
	r := ""
	lS := len(s)
	sub := 2
	if numRows == 1 {
		sub--
	}
	for i := 0; i < numRows; i++ {
		for j := i; j < lS; j += numRows + numRows - sub {
			p := j + ((numRows - i - 1) * 2)
			r += string([]byte{s[j]})
			if i != 0 && i != numRows-1 && p < lS {
				r += string([]byte{s[p]})
			}
		}
	}
	return r
}

func main() {
	a := []data{
		{"PAYPALISHIRING", "PAHNAPLSIIGYIR", 3},
		{"PAYPALISHIRING", "PINALSIGYAHRPI", 4},
		{"A", "A", 1},
	}
	fail := false
	for _, v := range a {
		out := convert2(v.init, v.rows)
		fmt.Println(out + " " + v.outp)
		if out != v.outp {
			fail = true
		}
	}
	if fail {
		fmt.Println("FAIL")
	} else {
		fmt.Println("SUCCESS")
	}
}
