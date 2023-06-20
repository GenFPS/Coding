package main

func removeDuplicates(nums []int) int {
	j := 1
	prev := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != prev {
			nums[j] = nums[i]
			j++
		}
		prev = nums[i]
	}
	return j
}
