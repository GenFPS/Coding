package main

import (
	"fmt"
	"sort"
	"strconv"
)

func perm(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	var n int
	fmt.Scan(&n)

	numbs := make([]int, 3)

	for i := 0; i < 3; i++ {
		if i == 0 {
			numbs[i] = n / 100
		}
		if i == 1 {
			numbs[i] = (n % 100) / 10
		}
		if i == 2 {
			numbs[i] = n % 10
		}
	}
	numbs2d := perm(numbs)
	numbsStr := make([]string, len(numbs2d))

	for i := range numbs2d {
		numS := ""
		for j := range numbs2d[i] {
			numS += strconv.Itoa(numbs2d[i][j])
		}
		numbsStr[i] = numS
	}
	newNumInt := make([]int, len(numbsStr))

	for i := range numbsStr {
		el, err := strconv.Atoi(numbsStr[i])
		if err != nil {
			fmt.Println(err)
		}
		newNumInt[i] = el
	}

	sort.Ints(newNumInt)
	for i := range newNumInt {
		fmt.Println(newNumInt[i])
	}
}
