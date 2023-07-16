package main

import (
	"sort"
)

func sortedSquares(nums []int) []int {
	sqrArr := make([]int, len(nums))

	for i := range nums {
		sqrArr[i] = nums[i] * nums[i]
	}
	sort.Ints(sqrArr)
	return sqrArr
}
