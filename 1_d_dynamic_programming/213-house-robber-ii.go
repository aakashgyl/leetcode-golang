package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	return max(robMax(nums[1:]), robMax(nums[:len(nums)-1]))
}

func robMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	var robMaxTillIndex []int
	robMaxTillIndex = append(robMaxTillIndex, nums[0])               // robMaxTillIndex[0]
	robMaxTillIndex = append(robMaxTillIndex, max(nums[0], nums[1])) //robMaxTillIndex[1]

	for i := 2; i < len(nums); i++ {
		robMaxTillIndex = append(robMaxTillIndex, max(nums[i]+robMaxTillIndex[i-2], robMaxTillIndex[i-1]))
	}
	return robMaxTillIndex[len(nums)-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
