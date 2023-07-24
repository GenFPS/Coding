package main

// O(n**2)
func sortArrayByParity1(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j]%2 == 0 {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// O(n)
func sortArrayByParity2(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	j := 0
	for i := range nums {
		if nums[i]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}

	}
	return nums
}
