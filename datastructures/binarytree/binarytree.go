package binarytree

import (
	"math/rand"
	"fmt"
)

type TreeNode struct {
	Val int32
	Left *TreeNode
	Right *TreeNode
}

type TreeNodeValidationFunc func(*TreeNode) bool

// GetRandomTree generate a random binary tree. Returns the root node of the tree.
func GetRandomTree (maxNodes int32, shuffle bool) *TreeNode {
	root := new (TreeNode)
	root.Val = 0
	pool := make([]int32, maxNodes)
	for i := maxNodes - 1; i >= 0; i-- {
		pool[i] = i
	}
	if shuffle {
		rand.Shuffle(len(pool), func(i, j int) {
			pool[i], pool[j] = pool[j], pool[i]
		})
	}
	setRandomChilds(root, &maxNodes, pool)
	return root
}

func setRandomChilds (node *TreeNode, maxStack *int32, pool []int32) {
	*maxStack -= 1
	if *maxStack < 0 {
		return
	}
	node.Val = pool[*maxStack]

	hasLeft := rand.Float32() < 0.5
	hasRight := rand.Float32() < 0.5
	for !hasLeft && !hasRight {
		hasLeft = rand.Float32() < 0.5
		hasRight = rand.Float32() < 0.5
	}

	if hasLeft {
		node.Left = new (TreeNode)
		setRandomChilds(node.Left, maxStack, pool)
	}
	
	if hasRight {
		node.Right = new (TreeNode)
		setRandomChilds(node.Right, maxStack, pool)
	}
}
