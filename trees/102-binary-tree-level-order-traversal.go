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

func (q *Queue) Enqueue(node *TreeNode) {
	*q = append(*q, node)
}

func (q *Queue) Dequeue() *TreeNode {
	if len(*q) == 0 {
		return nil
	}

	front := (*q)[0]
	*q = (*q)[1:]
	return front
}

func levelOrder(root *TreeNode) [][]int {
	var output [][]int

	if root == nil {
		return [][]int{}
	}

	q := Queue{}
	q.Enqueue(root)

	for len(q) != 0 {
		levelCount := len(q)
		var levelData []int

		for levelCount > 0 {
			node := q.Dequeue()
			levelData = append(levelData, (*node).Val)

			if (*node).Left != nil {
				q.Enqueue((*node).Left)
			}
			if (*node).Right != nil {
				q.Enqueue((*node).Right)
			}
			levelCount--
		}

		output = append(output, levelData)
	}

	return output
}
