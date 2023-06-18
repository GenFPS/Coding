package codeforces

import (
	"fmt"
)

// создаем 2дМассив -- сложность O(n^2)

func d2Array(N int) [][]int {
	d2Arr := make([][]int, N)

	for i := range d2Arr {
		var n int // 2-ой параметр длина подмассива
		fmt.Scan(&n)
		d2Arr[i] = make([]int, n)

		for j := range d2Arr[i] {
			var el int
			fmt.Scan(&el)
			d2Arr[i][j] = el
		}
	}
	return d2Arr
}

// считаем кол-во отрицательных чисел
func countNegative(arr []int) int {
	var counterNegative int
	for i := range arr {
		if arr[i] == -1 {
			counterNegative++
		}
	}
	return counterNegative
}

func main() {
	var N int // 1-ый параметр длина 2д массива
	fmt.Scan(&N)

	d2Arr := d2Array(N)

	for i := range d2Arr {
		negative := countNegative(d2Arr[i])

		var oper int

		if negative%2 == 1 {
			negative--
			oper++
		}
		for negative*2 > len(d2Arr[i]) {
			negative -= 2
			oper += 2
		}
		fmt.Printf("%d\n", oper)

	}
}
