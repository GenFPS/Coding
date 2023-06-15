package main

import (
	"fmt"
)

func reverseS(s []byte) {
	var j int
	for i := len(s) - 1; i >= len(s)/2; i-- {
		s[j], s[i] = s[i], s[j]
		j++
	}
	fmt.Println(s)
}

// func main() {
// 	s := "Hannah"
// 	bs := []byte(s)

// 	reverseS(bs)
// }
