package main

func singleNumber(nums []int) int {
	num := 0

	for _, val := range nums {
		num = num ^ val
	}

	return num
}
