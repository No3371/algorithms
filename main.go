package main

import (
	"bastudio/algorithms/bfs"
	"bastudio/algorithms/datastructures/binarytree"
	"bastudio/algorithms/dfs"
	"bastudio/algorithms/rot2d"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	switch os.Args[2] {
	case "bfs":
		{
			doBFS()
			break
		}
	case "dfs":
		{
			doDFS()
			break
		}
	case "rot2d":
		{
			doRot2D()
			break
		}
	}
}

func doBFS() {
	fmt.Println("Running BFS.")
	root := binarytree.GetRandomTree(22, false)
	bfs.BFS(root, func(r *binarytree.TreeNode) bool {
		fmt.Println(r.Val)
		if r.Val == 16 {
			return true
		} else {
			return false
		}
	})
}

func doDFS() {
	fmt.Println("Running DFS.")
	root := binarytree.GetRandomTree(22, false)
	dfs.DFS(root, func(r *binarytree.TreeNode) bool {
		fmt.Println(r.Val)
		if r.Val == 11 {
			return true
		} else {
			return false
		}
	})
}

func doRot2D () {
	fmt.Println("Running Rot2D")
	test := [][]int {
		{ 0, 1, 2, 3 },
		{ 4, 5, 6, 7 },
		{ 8, 9, 10, 11},
		{ 12, 13, 14, 15},
	}
	fmt.Printf("\n%v", test[0])
	fmt.Printf("\n%v", test[1])
	fmt.Printf("\n%v", test[2])
	fmt.Printf("\n%v", test[3])
	fmt.Printf("\n")
	rot2d.RotateCounterClockWise(&test)
	fmt.Printf("\n%v", test[0])
	fmt.Printf("\n%v", test[1])
	fmt.Printf("\n%v", test[2])
	fmt.Printf("\n%v", test[3])
}