package main

import "fmt"

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))

	// prefix filling
	prefixProduct := 1
	for i := 0; i < len(nums); i++ {
		output[i] = prefixProduct
		prefixProduct = prefixProduct * nums[i]
	}

	// postfix filling
	postfixProduct := 1
	for i := len(nums) - 1; i >= 0; i-- {
		output[i] = postfixProduct * output[i]
		postfixProduct = postfixProduct * nums[i]
	}

	return output
}
