package main

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	m := make(map[string]int)
	for _, v := range s {
		m[string(v)]++
	}
	if len(m) == 1 {
		return 1
	}
	if uniqSymb(s) {
		return len(s)
	} else {
		maxLen := 0
		maxSoFar := string(s[0])
		maxEndHere := ""

		for i := 0; i < len(s); i++ {
			maxEndHere = maxEndHere + string(s[i])

			if !uniqSymb(maxEndHere) {
				maxLen = maxOfLen(maxEndHere, maxSoFar)
				maxSoFar = maxEndHere
				maxEndHere = ""
			}
		}
		return maxLen
	}
}

func uniqSymb(s string) bool {
	m := make(map[string]int)

	for _, v := range s {
		m[string(v)]++
	}
	for k := range m {
		if m[k] > 1 {
			return false
		}
	}
	return true
}

func maxOfLen(s1, s2 string) int {
	if s1 > s2 {
		return len(s1)
	}
	return len(s2)
}
