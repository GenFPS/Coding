package main

import (
	"fmt"
	"strconv"
)

func addDigits(num int) int {
	numS := strconv.Itoa(num)
	var sumInt int

	for i := range numS {

		elInt, err := strconv.Atoi(string(numS[i]))
		if err != nil {
			fmt.Println("err")
		}
		sumInt += elInt
	}
	for sumInt >= 10 {
		num = sumInt
		sumInt = 0
		numS = strconv.Itoa(num)
		for i := range numS {
			elInt, err := strconv.Atoi(string(numS[i]))
			if err != nil {
				fmt.Println("err")
			}
			sumInt += elInt
		}
	}
	return sumInt
}
