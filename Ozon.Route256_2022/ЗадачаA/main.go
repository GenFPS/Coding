package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	nS := strconv.Itoa(n)
	nS = nS[1:4]
	newN, err := strconv.Atoi(nS)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newN)
}
