package main

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		fmt.Println(nums1)
	}
	if m == 0 {
		nums1 = append(nums1, nums2...)
		counter := 0
		sort.Ints(nums1)
		for i := range nums1 {
			if nums1[i] == 0 {
				counter++
			}
		}
		fmt.Println(nums1[counter:])
	}
	if n != 0 && m != 0 {
		nums1 = append(nums1, nums2...)
		counter := 0
		sort.Ints(nums1)
		for i := range nums1 {
			if nums1[i] == 0 {
				counter++
			}
		}
		fmt.Println(nums1[counter:])
	}
}
