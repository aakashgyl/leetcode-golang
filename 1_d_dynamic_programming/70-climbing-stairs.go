package main

func climbStairs(n int) int {
	data := make(map[int]int)

	data[0] = 0
	data[1] = 1
	data[2] = 2

	for i := 3; i <= n; i++ {
		data[i] = data[i-1] + data[i-2]
	}

	return data[n]
}
