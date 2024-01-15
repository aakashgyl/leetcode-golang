package main

func longestConsecutive(nums []int) int {
	dataMap := make(map[int]bool)
	for _, num := range nums {
		dataMap[num] = true
	}

	var maxCount int

	for _, num := range nums {
		if dataMap[num-1] == false { // number doesnt exist, so num is start of sequence
			var currCount int = 1
			// check sequence length
			for i := 1; i < len(nums); i++ {
				if dataMap[num+i] == true {
					currCount++
				} else {
					break
				}
			}

			maxCount = getMax(maxCount, currCount)
		}
	}

	return maxCount
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
