package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var count = 0

func kthSmallest(root *TreeNode, k int) int {
	count = 0
	node := inorder(root, k)
	if node != nil {
		return node.Val
	}
	return -1
}

func inorder(root *TreeNode, k int) *TreeNode {
	fmt.Printf("iter: %d\n", count)
	if root == nil {
		return nil
	}

	// search in left subtree first
	target := inorder(root.Left, k)
	if target != nil {
		return target
	}

	// root is the element being searched
	count++
	if count == k {
		return root
	}

	// search in right subtree
	return inorder(root.Right, k)
}
