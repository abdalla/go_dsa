package main

import (
	"fmt"

	"github.com/abdalla/go_dsa/tree"
)

// - ROOT > LEFT > RIGHT
func preorder(t *tree.Tree) (traversal []int) {
	if t == nil {
		return
	}

	traversal = append(traversal, t.Value)
	traversal = append(traversal, preorder(t.Left)...)
	traversal = append(traversal, preorder(t.Right)...)

	return
}

// - LEFT > ROOT > RIGHT
func inorder(t *tree.Tree) (traversal []int) {
	if t == nil {
		return
	}

	traversal = append(traversal, inorder(t.Left)...)
	traversal = append(traversal, t.Value)
	traversal = append(traversal, inorder(t.Right)...)

	return
}

// - LEFT > RIGHT > ROOT
func postorder(t *tree.Tree) (traversal []int) {
	if t == nil {
		return
	}

	traversal = append(traversal, postorder(t.Left)...)
	traversal = append(traversal, postorder(t.Right)...)
	traversal = append(traversal, t.Value)

	return
}

func main() {
	t := tree.CreateBinaryTree()
	fmt.Printf("%v \n", preorder(t))
	fmt.Printf("%v \n", inorder(t))
	fmt.Printf("%v \n", postorder(t))
}
