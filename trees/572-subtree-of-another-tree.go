package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil { // a nil subtree is always a subtree of nil or non-nil tree
		return true
	}

	if root == nil { // if main tree is nil but subtree is not nil, its false
		return false
	}

	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil || root.Val != subRoot.Val {
		return false
	}

	return isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
}
