package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	var cutters int
	pieces := 1

	for pieces < N {
		pieces *= 2
		cutters++
	}
	fmt.Println(cutters)
}
