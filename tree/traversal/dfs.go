package traversal

import (
	"fmt"

	"github.com/abdalla/go_dsa/tree"
)

// - ROOT > LEFT > RIGHT
func Preorder(t *tree.Tree) (traversal []int) {
	if t == nil {
		return
	}

	traversal = append(traversal, t.Value)
	traversal = append(traversal, Preorder(t.Left)...)
	traversal = append(traversal, Preorder(t.Right)...)

	return
}

// - LEFT > ROOT > RIGHT
func Inorder(t *tree.Tree) (traversal []int) {
	if t == nil {
		return
	}

	traversal = append(traversal, Inorder(t.Left)...)
	traversal = append(traversal, t.Value)
	traversal = append(traversal, Inorder(t.Right)...)

	return
}

// - LEFT > RIGHT > ROOT
func Postorder(t *tree.Tree) (traversal []int) {
	if t == nil {
		return
	}

	traversal = append(traversal, Postorder(t.Left)...)
	traversal = append(traversal, Postorder(t.Right)...)
	traversal = append(traversal, t.Value)

	return
}

func dfs() {
	t := tree.CreateBinaryTree()
	fmt.Printf("%v \n", Preorder(t))
	fmt.Printf("%v \n", Inorder(t))
	fmt.Printf("%v \n", Postorder(t))
}
