package main

import "container/heap"

type MaxHeap []int

func findKthLargest(nums []int, k int) int {
	maxHeap := MaxHeap(nums)
	heap.Init(&maxHeap)

	res := 0
	for i := 0; i < k; i++ {
		res = heap.Pop(&maxHeap).(int)
	}
	return res
}

func (m MaxHeap) Len() int           { return len(m) }
func (m MaxHeap) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m MaxHeap) Less(i, j int) bool { return m[i] > m[j] }

func (m *MaxHeap) Pop() any {
	n := len(*m) - 1
	popVal := (*m)[n]
	*m = (*m)[:n]
	return popVal
}

func (m *MaxHeap) Push(val any) {
	*m = append(*m, val.(int))
}
