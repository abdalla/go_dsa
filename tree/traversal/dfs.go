package main

import (
	"fmt"

	"github.com/abdalla/go_dsa/tree"
)

func preorder(t *tree.Tree) (traversal []int) {
	if t == nil {
		return
	}

	traversal = append(traversal, t.Value)
	traversal = append(traversal, preorder(t.Left)...)
	traversal = append(traversal, preorder(t.Right)...)

	return
}

func main() {
	t := tree.CreateBinaryTree()
	fmt.Printf("%v", preorder(t))
}
