package main

func maxArea(height []int) int {
	l := 0
	r := len(height) - 1

	max := 0

	for l < r {
		area := (r - l) * getMin(height[l], height[r])
		max = getMax(max, area)

		if height[l] > height[r] {
			r--
		} else {
			l++
		}
	}

	return max
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
