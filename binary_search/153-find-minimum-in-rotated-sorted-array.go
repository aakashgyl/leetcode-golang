package main

func findMin(nums []int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		if nums[l] <= nums[r] {
			return nums[l]
		}

		mid := (l + r) / 2
		if mid > 0 && nums[mid] < nums[mid-1] {
			return nums[mid]
		}

		if nums[mid] < nums[r] { // clean array
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return nums[l]
}
