package main

import (
	"sort"
)

func replaceElements(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		if i == len(arr)-1 {
			arr[i] = -1
		}
		subArr := make([]int, len(arr)-(i+1))
		if len(subArr) == 0 {
			continue
		} else {
			copy(subArr, arr[i+1:])
			sort.Ints(subArr)
			arr[i] = subArr[len(subArr)-1]
		}

	}
	return arr
}
