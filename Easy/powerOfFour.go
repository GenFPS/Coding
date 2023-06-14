package main

import (
	"math"
)

func isPowerOfFour(n int) bool {
	var pFour float64
	var i float64
	for int(pFour) <= n {
		pFour = math.Pow(4, i)

		if int(pFour) == n {
			return true
		}
		i++
	}
	return false
}
