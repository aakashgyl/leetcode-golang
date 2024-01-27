package main

func missingNumber(nums []int) int {
	val := 0
	i, num := 0, 0

	for i, num = range nums {
		val = val ^ num ^ i
	}
	return val ^ (i + 1)
}
