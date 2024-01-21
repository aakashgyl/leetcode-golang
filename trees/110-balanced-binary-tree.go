package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if absDiffInt(leftHeight, rightHeight) > 1 {
		return false
	}

	return isBalanced(root.Left) && isBalanced(root.Right)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + getMax(height(root.Left), height(root.Right))
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
