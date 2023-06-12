package main

func moveZeroes(nums []int) {
	var nums1, zeros []int
	for i := range nums {
		if nums[i] != 0 {
			nums1 = append(nums1, nums[i])
		} else if nums[i] == 0 {
			zeros = append(zeros, nums[i])
		}
	}
	nums1 = append(nums1, zeros...)

	for j := range nums1 {
		nums[j] = nums1[j]
	}
}
