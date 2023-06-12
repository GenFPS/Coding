package main

func majorityElement(nums []int) int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v]++
	}

	var maxK, maxV int
	for k, v := range m {
		if v > maxV {
			maxV = v
			maxK = k
		}
	}
	return maxK
}
