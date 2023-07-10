package main

import (
	"strconv"
)

func hammingWeight(num uint32) int {
	sNum := strconv.FormatUint(uint64(num), 2)
	counterBits := 0
	for i := range sNum {
		if string(sNum[i]) == "1" {
			counterBits++
		}
	}
	return counterBits
}
