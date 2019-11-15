package dfs

import (
	. "bastudio/algorithms/datastructures/binarytree"
)

func DFS (root *TreeNode, validator TreeNodeValidationFunc) *TreeNode{
	var resultPointer = new(TreeNode)
	if recursive(root, resultPointer, validator) {
		return resultPointer
	} else {
		return nil
	}
}

func recursive (root *TreeNode, resultPointer *TreeNode, validator TreeNodeValidationFunc) bool {
	if validator(root) {
		return true
	}
	if root.Left != nil {
		if recursive(root.Left, resultPointer, validator) {
			return true
		}
	} else {
		if root.Right != nil {
			if recursive(root.Right, resultPointer, validator) {
				return true
			}
		}
	}
	return false
}