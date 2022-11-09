package tree

import (
	"github.com/abdalla/go_dsa/tree"
)

func BstFromPreorder(preorder []int) *tree.Tree {
	root := &tree.Tree{
		Value: preorder[0],
	}

	stack := []*tree.Tree{root}

	for i := 1; i < len(preorder); i++ {
		lastOnStack := stack[len(stack)-1]
		if preorder[i] < lastOnStack.Value {
			node := &tree.Tree{
				Value: preorder[i],
			}

			lastOnStack.Left = node
			stack = append(stack, node)
		} else {
			for len(stack) > 0 && stack[len(stack)-1].Value < preorder[i] {
				lastOnStack = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}

			node := &tree.Tree{
				Value: preorder[i],
			}
			lastOnStack.Right = node
			stack = append(stack, node)
		}
	}

	return root
}
