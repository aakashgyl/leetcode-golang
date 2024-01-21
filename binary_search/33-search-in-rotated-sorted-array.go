package main

func search(nums []int, target int) int {
	l := 0
	r := len(nums) - 1

	for l <= r {
		mid := (l + r) / 2

		if target == nums[mid] {
			return mid
		}

		if nums[l] <= nums[mid] { // left array is sorted
			if target >= nums[l] && target <= nums[mid] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else { // right array is sorted
			if target >= nums[mid] && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
