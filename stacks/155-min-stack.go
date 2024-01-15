package main

type MinStack struct {
	data    []int
	minData []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)

	length := len(this.minData)
	minVal := val
	if length > 0 {
		minVal = getMin(this.minData[len(this.minData)-1], val)
	}
	this.minData = append(this.minData, minVal)
}

func (this *MinStack) Pop() {
	length := len(this.minData)
	if length < 0 {
		return
	}

	this.data = this.data[:length-1]
	this.minData = this.minData[:length-1]
}

func (this *MinStack) Top() int {
	length := len(this.data)
	if length < 0 {
		return 0
	}
	return this.data[length-1]
}

func (this *MinStack) GetMin() int {
	length := len(this.minData)
	if length < 0 {
		return 0
	}
	return this.minData[length-1]
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
