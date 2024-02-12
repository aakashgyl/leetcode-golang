package main

import "container/heap"

type MinHeap [][]int

func kClosest(points [][]int, k int) [][]int {
	var res [][]int
	minHeap := MinHeap(points)
	heap.Init(&minHeap)

	for i := 0; i < k; i++ {
		res = append(res, (heap.Pop(&minHeap)).([]int))
	}
	return res
}

func (m MinHeap) Len() int      { return len(m) }
func (m MinHeap) Swap(i, j int) { m[i], m[j] = m[j], m[i] }

func (m MinHeap) Less(i, j int) bool {
	di := m[i][0]*m[i][0] + m[i][1]*m[i][1]
	dj := m[j][0]*m[j][0] + m[j][1]*m[j][1]

	return di < dj
}

func (m *MinHeap) Pop() any {
	n := len(*m) - 1
	popVal := (*m)[n]
	*m = (*m)[:n]
	return popVal
}

func (m *MinHeap) Push(val any) {
	*m = append(*m, val.([]int))
}
