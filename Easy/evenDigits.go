package main

import (
	"strconv"
)

func findNumbers(nums []int) int {
	counter := 0

	for i := 0; i < len(nums); i++ {
		numS := strconv.Itoa(nums[i])
		if len(numS)%2 == 0 {
			counter++
		}
	}
	return counter
}
