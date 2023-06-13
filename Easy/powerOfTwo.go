package main

import (
	"math"
)

func isPowerOfTwo(n int) bool {
	var pTwo float64
	var i float64
	for int(pTwo) <= n {
		pTwo = math.Pow(2, i)

		if int(pTwo) == n {
			return true
		}
		i++
	}
	return false
}
