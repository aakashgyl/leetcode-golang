package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var data []int
	inorder(root, &data)

	prev := data[0]
	for i := 1; i < len(data); i++ {
		if prev >= data[i] {
			return false
		}
		prev = data[i]
	}
	return true
}

func inorder(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, data)
	*data = append(*data, root.Val)
	inorder(root.Right, data)
}
