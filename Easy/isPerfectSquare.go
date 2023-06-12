package main

import "math"

func isPerfectSquare(num int) bool {
	if num == 1 {
		return true
	} else {
		var numFloat float64 = 2
		for {
			float := math.Pow(numFloat, 2)
			if float == float64(num) {
				return true
			}
			if float > float64(num) {
				return false
			}
			numFloat++
		}
	}
}
