package tree

import "github.com/abdalla/go_dsa/tree"

func InvertBT(root *tree.Tree) *tree.Tree {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	InvertBT(root.Left)
	InvertBT(root.Right)

	return root
}
