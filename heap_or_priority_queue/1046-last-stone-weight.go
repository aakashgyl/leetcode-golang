package main

import "container/heap"

type maxHeap []int

func lastStoneWeight(stones []int) int {
	maxheap := maxHeap(stones)
	heap.Init(&maxheap)

	for maxheap.Len() > 1 {
		s1 := heap.Pop(&maxheap).(int)
		s2 := heap.Pop(&maxheap).(int)

		if s1 != s2 {
			heap.Push(&maxheap, s1-s2)
		}
	}

	res := 0
	if maxheap.Len() == 1 {
		res = heap.Pop(&maxheap).(int)
	}

	return res
}

func (m maxHeap) Len() int {
	return len(m)
}

func (m maxHeap) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func (m maxHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m *maxHeap) Push(val any) {
	*m = append(*m, val.(int))
}

func (m *maxHeap) Pop() any {
	n := len(*m) - 1
	popVal := (*m)[n]
	*m = (*m)[:n]
	return popVal
}
