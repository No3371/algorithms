package bfs

import (
	. "bastudio/algorithms/datastructures/binarytree"
)

func BFS (root *TreeNode, validation TreeNodeValidationFunc) *TreeNode {
	open := make([]*TreeNode, 0)
		
	anchor := 0
	open = append(open, root)
	for anchor < len(open) {
		if validation(open[anchor]) {
			return open[anchor]
		}
		if (open[anchor].Left != nil) {
			open = append(open, open[anchor].Left)
		}
		if (open[anchor].Right != nil) {
			open = append(open, open[anchor].Right)
		}
		anchor++
	}
	return nil
}