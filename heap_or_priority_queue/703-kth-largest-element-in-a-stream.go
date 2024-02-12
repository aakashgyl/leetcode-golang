package main

import "math"

type KthLargest struct {
	data []int
	size int
}

func Constructor(k int, nums []int) KthLargest {
	heap := KthLargest{
		data: make([]int, k),
		size: k,
	}

	for i := 0; i < k; i++ {
		heap.data[i] = math.MinInt
	}

	// heapify first k elements
	copy(heap.data, nums)
	for i := k/2 - 1; i >= 0; i-- {
		heap.minHeapify(i)
	}

	// if nums had more than k elements, add them one by one
	if len(nums) > k {
		for i := k; i < len(nums); i++ {
			heap.Add(nums[i])
		}
	}
	return heap
}

func (this *KthLargest) minHeapify(i int) {
	smallest := i
	lchild := 2*i + 1
	rchild := 2*i + 2

	if lchild < this.size && this.data[smallest] > this.data[lchild] {
		smallest = lchild
	}

	if rchild < this.size && this.data[smallest] > this.data[rchild] {
		smallest = rchild
	}

	if smallest != i {
		this.data[i], this.data[smallest] = this.data[smallest], this.data[i]
		this.minHeapify(smallest)
	}
}

func (this *KthLargest) Add(val int) int {
	if this.data[this.size-1] == math.MinInt {
		this.data[this.size-1] = val
		this.minHeapifyBottomUp(this.size - 1)
	} else if val > this.data[0] {
		this.data[0] = val
		this.minHeapify(0)
	}

	return this.data[0]
}

func (this *KthLargest) minHeapifyBottomUp(i int) {
	parent := (i - 1) / 2

	if parent >= 0 {
		if this.data[parent] > this.data[i] {
			this.data[parent], this.data[i] = this.data[i], this.data[parent]
			this.minHeapifyBottomUp(parent)
		}
	}
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
