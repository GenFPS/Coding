package main

// import "fmt"

func strStr(haystack string, needle string) int {
	h := len(haystack)
	n := len(needle)

	if needle == "" {
		return 0
	} else {
		for i := 0; i <= h-n; i++ {
			if haystack[i:i+n] == needle {
				return i
			}
		}
	}
	return -1
}

// func main() {
// 	haystack := "sadbutsad"
// 	needle := "sad"

// 	fmt.Println(strStr(haystack, needle))
// }
