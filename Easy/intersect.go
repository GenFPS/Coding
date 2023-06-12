package main

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v]++
	}

	var nums3 []int

	for _, n := range nums2 {
		if v, ok := m[n]; ok {
			if v != 0 {
				nums3 = append(nums3, n)
				m[n] = v - 1
			}
		}
	}
	return nums3
}
