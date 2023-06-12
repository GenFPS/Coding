package main

import (
	"math"
)

func isPowerOfThree(n int) bool {
	var d float64
	n1 := math.Pow(3, d)

	if n <= 0 {
		return false
	}
	if n > 0 {
		for int(n1) <= n {
			n1 = math.Pow(3, d)

			if int(n1) == n {
				return true
			}
			if int(n1) > n {
				return false
			}
			d++
		}
	}
	return true
}
