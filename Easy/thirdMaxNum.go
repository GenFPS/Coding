package main

import (
	"sort"
)

func thirdMax(nums []int) int {
	sort.Ints(nums)
	if len(nums) <= 2 {
		return nums[len(nums)-1]
	}
	// remove duplicates from sorted arr
	inResult := make(map[int]bool)
	var numsFilt []int
	for _, num := range nums {
		if _, ok := inResult[num]; !ok {
			inResult[num] = true
			numsFilt = append(numsFilt, num)
		}
	}
	if len(numsFilt) <= 2 {
		return numsFilt[len(numsFilt)-1]
	}
	return numsFilt[len(numsFilt)-3]
}
