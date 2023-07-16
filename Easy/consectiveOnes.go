package main

import (
	"sort"
)

func findMaxConsecutiveOnes(nums []int) int {
	var consecArr []int

	counter := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter++
		}
		if nums[i] == 0 {
			consecArr = append(consecArr, counter)
			counter = 0
		}
		if i == len(nums)-1 {
			consecArr = append(consecArr, counter)
		}
	}
	sort.Ints(consecArr)
	return consecArr[len(consecArr)-1]
}
