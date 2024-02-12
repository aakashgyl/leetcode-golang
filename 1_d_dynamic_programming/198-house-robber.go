package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	robMap := make(map[int]int) // index - max till that index
	robMap[0] = nums[0]
	robMap[1] = getMax(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		robMap[i] = getMax(nums[i]+robMap[i-2], robMap[i-1])
	}
	return robMap[len(nums)-1]
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
