package main

import (
	"bastudio/algorithms/bfs"
	"bastudio/algorithms/dfs"
	"bastudio/algorithms/datastructures/binarytree"
	"fmt"
	"os"
	"math/rand"
	"time"
)

func main () {
	rand.Seed(time.Now().UnixNano())
	switch os.Args[2] {
		case "bfs": {
			doBFS()
			break
		}
		case "dfs": {
			doDFS()
			break
		}
	}
}

func doBFS () {
	fmt.Println("Running BFS.")
	root := binarytree.GetRandomTree(22, false)
	bfs.BFS(root, func (r *binarytree.TreeNode) bool {
		fmt.Println(r.Val)
		if r.Val == 16 {
			return true
		} else {
			return false
		}
	})
}

func doDFS () {
	fmt.Println("Running DFS.")
	root := binarytree.GetRandomTree(22, false)
	dfs.DFS(root, func (r *binarytree.TreeNode) bool {
		fmt.Println(r.Val)
		if r.Val == 11 {
			return true
		} else {
			return false
		}
	})
	
}