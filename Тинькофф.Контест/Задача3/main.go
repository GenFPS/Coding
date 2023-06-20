package main

import (
	"fmt"
)

func make1DArr(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		var el int
		fmt.Scan(&el)
		arr[i] = el
	}
	return arr
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	var n, t int = 5, 5
	fmt.Scan(&n, &t)

	arr := make1DArr(n)

	var idx int
	fmt.Scan(&idx)

	countMinutes := 0
	if t >= arr[idx-1] {
		for i := 0; i < len(arr); i++ {
			if i == len(arr)-1 {
				break
			}
			countMinutes += arr[i] - arr[i+1]
		}
	}
	if t < arr[idx-1] {
		countMinutes += arr[0] - arr[idx-1]
		newArr := removeIndex(arr, idx-1)

		for i := 0; i < len(newArr); i++ {
			if i == len(newArr)-1 {
				break
			}
			countMinutes += newArr[i] - newArr[i+1]
		}
	}
	countMinutes += (-1 * countMinutes) * 2
	fmt.Println(countMinutes)
}
