package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}

func topKFrequent(nums []int, k int) []int {
	var output []int

	// create map with count of each element
	// nums = [1,1,1,2,2,3], k = 2
	// countMap -> {{1,3}, {2,2}, {3,1}} (num, count)
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num] = countMap[num] + 1
	}

	// create array with count as index and value as list of num with that count
	// array = 1->3, 2->2, 3->1 (count->num)
	array := make([][]int, len(nums)+1) // slice length cannot be more than unique elements in nums
	for num, count := range countMap {
		array[count] = append(array[count], num)
	}

	// find top k elements
	count := 0
	for i := len(nums); i > 0; i-- {
		if len(array[i]) != 0 {
			for _, num := range array[i] {
				output = append(output, num)
				count = count + 1
				if count == k {
					return output
				}
			}
		}
	}

	return output
}
