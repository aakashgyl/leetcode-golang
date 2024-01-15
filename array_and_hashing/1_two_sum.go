package main

func twoSum(nums []int, target int) []int {
	dataMap := make(map[int]int, len(nums))

	for i, val := range nums {
		otherNum := target - val
		if j, ok := dataMap[otherNum]; ok {
			return []int{i, j}
		}
		dataMap[val] = i
	}
	return []int{}
}
