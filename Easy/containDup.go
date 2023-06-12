package main

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k := range m {
		if m[k] > 1 {
			return true
		}
	}
	return false
}
