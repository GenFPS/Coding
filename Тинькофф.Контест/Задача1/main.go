package main

import (
	"fmt"
)

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)

	if B >= D {
		fmt.Println(A)
	} else {
		formula := A + (D-B)*C
		fmt.Println(formula)
	}
}
