package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: nil,
	}
	fmt.Println(diameterOfBinaryTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var max int

func diameterOfBinaryTree(root *TreeNode) int {
	return preorder(root)
}

func preorder(root *TreeNode) int {
	if root == nil {
		return 0
	}
	diaroot := getDia(root)
	preorder(root.Left)
	preorder(root.Right)
	max = getMax(diaroot, max)
	return max
}

func getDia(root *TreeNode) int {
	l := getHeight(root.Left)
	r := getHeight(root.Right)
	return l + r
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)

	return 1 + getMax(leftHeight, rightHeight)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
