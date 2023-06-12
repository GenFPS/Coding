package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(x int) int {
	s1 := strconv.Itoa(x)
	s2 := strings.Split(s1, "")

	if s2[0] != "-" && s2[len(s2)-1] != "0" { // case x = 123
		reverseString(s2)
	} else if s2[0] == "-" && s2[len(s2)-1] != "0" { // case x = -123
		var s3 []string

		s3 = append(s3, s2[0])
		s2 = s2[1:]

		reverseString(s2)
		s2 = append(s3, s2...)

	} else if s2[0] == "-" && s2[len(s2)-1] == "0" { // case x = -1230
		var s3 []string

		s3 = append(s3, s2[0])
		s2 = s2[1 : len(s2)-1]
		reverseString(s2)

		s2 = append(s3, s2...)

	} else if s2[0] != "-" && s2[len(s2)-1] == "0" { // case x = 120
		s2 = s2[:len(s2)-1]

		reverseString(s2)
	}

	var s3 string
	for i := range s2 {
		s3 += s2[i]
	}

	N2, err := strconv.Atoi(s3)
	if err != nil {
		fmt.Println(err)
	}

	if N2 < -2147483648 || N2 > 2147483647 {
		return 0
	}
	return N2
}

func reverseString(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
