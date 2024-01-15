package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Queue []*TreeNode

func (q *Queue) Push(val *TreeNode) {
	*q = append(*q, val)
}

func (q *Queue) Pop() *TreeNode {
	if len(*q) == 0 {
		return nil
	}
	top := (*q)[0]
	*q = (*q)[1:]
	return top
}

func levelOrder(root *TreeNode) [][]int {
	var output [][]int

	if root == nil {
		return [][]int{}
	}

	q := Queue{}
	q.Push(root)

	for len(q) != 0 {
		levelCount := len(q)
		var levelData []int

		for levelCount > 0 {
			node := q.Pop()
			levelData = append(levelData, (*node).Val)

			if (*node).Left != nil {
				q.Push((*node).Left)
			}
			if (*node).Right != nil {
				q.Push((*node).Right)
			}
			levelCount--
		}

		output = append(output, levelData)
	}

	return output
}
